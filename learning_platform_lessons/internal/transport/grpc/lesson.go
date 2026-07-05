package grpc

import (
	"context"
	lessonGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/lesson"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"learning-platform/lessons/internal/dto"
	"learning-platform/lessons/internal/models/enum"
)

type LessonGRPCServer struct {
	lessonGRPC.UnimplementedLessonServer
	LessonService LessonService
	logger        *zap.Logger
}

type LessonService interface {
	GetOneLesson(lessonID int64) (*dto.LessonResponse, error)
	GetLessonsByUserId(userID int64) ([]dto.LessonResponse, error)
	CreateLesson(lesson dto.CreateLesson) (*dto.LessonResponse, error)
	UpdateLesson(lesson dto.UpdateLesson) (*dto.LessonResponse, error)
	UpdateLessonStatus(lessonID int64, status enum.LessonStatus) error
	GetLessonsByTutorId(tutorID int64) ([]dto.LessonResponse, error)
}

func NewLessonGRPCServer(
	logger *zap.Logger,
	lesson LessonService,
) lessonGRPC.LessonServer {
	return &LessonGRPCServer{
		logger:        logger,
		LessonService: lesson,
	}
}

func (l *LessonGRPCServer) GetOneLesson(
	ctx context.Context,
	in *lessonGRPC.GetOneLessonRequest,
) (*lessonGRPC.GetOneLessonResponse, error) {
	res, err := l.LessonService.GetOneLesson(in.GetId())
	if err != nil {
		l.logger.Error(
			"failed get lesson by id",
			zap.Int64("lessonID", in.GetId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed get lesson by id")
	}

	var resLessonMedias []*lessonGRPC.MediaItem
	for _, oneMedia := range res.MediaItems {
		resLessonMedias = append(
			resLessonMedias,
			&lessonGRPC.MediaItem{
				Id:        oneMedia.ID,
				LessonId:  oneMedia.LessonID,
				S3Link:    oneMedia.S3Link,
				S3Preview: oneMedia.S3Preview,
				Type:      enumToProtoType(oneMedia.Type),
			},
		)
	}

	return &lessonGRPC.GetOneLessonResponse{
		Id:         res.ID,
		BoardId:    res.BoardID,
		MeetLink:   res.MeetLink,
		StartTime:  timestamppb.New(res.StartTime),
		MediaItems: resLessonMedias,
		Duration:   res.Duration,
		TutorId:    res.TutorID,
		UserIds:    res.UserIDs,
		Status:     enumToProtoStatus(res.Status),
	}, nil
}

func (l *LessonGRPCServer) GetLessonsByUserId(
	ctx context.Context,
	in *lessonGRPC.GetLessonsByUserIdRequest,
) (*lessonGRPC.GetLessonsByUserIdResponse, error) {
	res, err := l.LessonService.GetLessonsByUserId(in.GetUserId())
	if err != nil {
		l.logger.Error(
			"failed get lesson by user id",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
	}

	var resLessons []*lessonGRPC.GetOneLessonResponse
	for _, oneLesson := range res {
		resLessons = append(resLessons, dtoToGRPCResponse(oneLesson))
	}

	return &lessonGRPC.GetLessonsByUserIdResponse{
		Lessons: resLessons,
	}, nil
}

func (l *LessonGRPCServer) CreateLesson(
	ctx context.Context,
	in *lessonGRPC.CreateLessonRequest,
) (*lessonGRPC.CreateLessonResponse, error) {
	var mediaItems []dto.MediaItem

	for _, oneMedia := range in.GetMediaItems() {
		mediaItems = append(
			mediaItems,
			dto.MediaItem{
				S3Link:    oneMedia.S3Link,
				S3Preview: oneMedia.S3Preview,
				Type:      protoToEnumType(oneMedia.Type),
			},
		)
	}

	lesson := dto.CreateLesson{
		BoardID:    in.BoardId,
		MeetLink:   in.MeetLink,
		StartTime:  in.GetStartTime().AsTime(),
		Duration:   in.GetDuration(),
		TutorID:    in.GetTutorId(),
		MediaItems: mediaItems,
		UserIDs:    in.GetUserIds(),
	}
	res, err := l.LessonService.CreateLesson(lesson)
	if err != nil {
		l.logger.Error(
			"failed to create lesson",
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to create lesson")
	}

	var resLessonMedias []*lessonGRPC.MediaItem
	for _, oneMedia := range res.MediaItems {
		resLessonMedias = append(
			resLessonMedias,
			&lessonGRPC.MediaItem{
				Id:        oneMedia.ID,
				LessonId:  oneMedia.LessonID,
				S3Link:    oneMedia.S3Link,
				S3Preview: oneMedia.S3Preview,
				Type:      enumToProtoType(oneMedia.Type),
			},
		)
	}

	return &lessonGRPC.CreateLessonResponse{
		Lesson: dtoToGRPCResponse(*res),
	}, nil
}

func (l *LessonGRPCServer) UpdateLesson(
	ctx context.Context,
	in *lessonGRPC.UpdateLessonRequest,
) (*lessonGRPC.UpdateLessonResponse, error) {
	var mediaItems []dto.MediaItem

	for _, oneMedia := range in.GetMediaItems() {
		mediaItems = append(
			mediaItems,
			dto.MediaItem{
				S3Link:    oneMedia.S3Link,
				S3Preview: oneMedia.S3Preview,
				Type:      protoToEnumType(oneMedia.Type),
			},
		)
	}

	lesson := dto.UpdateLesson{
		ID:              in.GetId(),
		BoardID:         in.BoardId,
		MeetLink:        in.MeetLink,
		StartTime:       in.GetStartTime().AsTime(),
		Duration:        in.GetDuration(),
		MediaItems:      mediaItems,
		UserIDs:         in.GetUserIds(),
		DeletedUserIDs:  in.GetDeletedUserIds(),
		DeletedMediaIDs: in.GetDeletedMediaIds(),
	}
	res, err := l.LessonService.UpdateLesson(lesson)
	if err != nil {
		l.logger.Error(
			"failed to create lesson",
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to create lesson")
	}

	var resLessonMedias []*lessonGRPC.MediaItem
	for _, oneMedia := range res.MediaItems {
		resLessonMedias = append(
			resLessonMedias,
			&lessonGRPC.MediaItem{
				Id:        oneMedia.ID,
				LessonId:  oneMedia.LessonID,
				S3Link:    oneMedia.S3Link,
				S3Preview: oneMedia.S3Preview,
				Type:      enumToProtoType(oneMedia.Type),
			},
		)
	}
	return &lessonGRPC.UpdateLessonResponse{
		Lesson: dtoToGRPCResponse(*res),
	}, nil
}

func (l *LessonGRPCServer) UpdateLessonStatus(
	ctx context.Context,
	in *lessonGRPC.UpdateLessonStatusRequest,
) (*lessonGRPC.UpdateLessonStatusResponse, error) {
	err := l.LessonService.UpdateLessonStatus(
		in.GetId(),
		protoToEnumStatus(in.GetStatus()),
	)
	if err != nil {
		l.logger.Error(
			"failed update status",
			zap.Int64("lessonID", in.GetId()),
			zap.String("status", string(protoToEnumStatus(in.GetStatus()))),
			zap.Error(err),
		)
	}

	return &lessonGRPC.UpdateLessonStatusResponse{}, nil
}

func (l *LessonGRPCServer) GetLessonsByTutorId(
	ctx context.Context,
	in *lessonGRPC.GetLessonsByTutorIdRequest,
) (*lessonGRPC.GetLessonsByTutorIdResponse, error) {
	res, err := l.LessonService.GetLessonsByTutorId(in.GetTutorId())
	if err != nil {
		l.logger.Error(
			"failed get lesson by tutor id",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
	}

	var resLessons []*lessonGRPC.GetOneLessonResponse
	for _, oneLesson := range res {
		resLessons = append(resLessons, dtoToGRPCResponse(oneLesson))
	}

	return &lessonGRPC.GetLessonsByTutorIdResponse{
		Lessons: resLessons,
	}, nil
}

func enumToProtoType(mediaType enum.MediaType) lessonGRPC.MediaType {
	switch mediaType {
	case enum.TypeImage:
		return lessonGRPC.MediaType_IMAGE
	case enum.TypeVideo:
		return lessonGRPC.MediaType_VIDEO
	default:
		return lessonGRPC.MediaType_MEDIA_TYPE_UNSPECIFIED
	}
}

func protoToEnumType(mediaType lessonGRPC.MediaType) enum.MediaType {
	switch mediaType {
	case lessonGRPC.MediaType_IMAGE:
		return enum.TypeImage
	case lessonGRPC.MediaType_VIDEO:
		return enum.TypeVideo
	default:
		return ""
	}
}

func protoToEnumStatus(lessonStatus lessonGRPC.LessonStatus) enum.LessonStatus {
	switch lessonStatus {
	case lessonGRPC.LessonStatus_SCHEDULED:
		return enum.StatusScheduled
	case lessonGRPC.LessonStatus_IN_PROCESS:
		return enum.StatusInProcess
	case lessonGRPC.LessonStatus_COMPLETED:
		return enum.StatusCompleted
	case lessonGRPC.LessonStatus_CANCELLED:
		return enum.StatusCancelled
	default:
		return ""
	}
}

func enumToProtoStatus(lessonStatus enum.LessonStatus) lessonGRPC.LessonStatus {
	switch lessonStatus {
	case enum.StatusScheduled:
		return lessonGRPC.LessonStatus_SCHEDULED
	case enum.StatusInProcess:
		return lessonGRPC.LessonStatus_IN_PROCESS
	case enum.StatusCompleted:
		return lessonGRPC.LessonStatus_COMPLETED
	case enum.StatusCancelled:
		return lessonGRPC.LessonStatus_CANCELLED
	default:
		return lessonGRPC.LessonStatus_LESSON_STATUS_UNSPECIFIED
	}
}

func dtoToGRPCResponse(lesson dto.LessonResponse) *lessonGRPC.GetOneLessonResponse {
	var lessonMedias []*lessonGRPC.MediaItem

	for _, oneMedia := range lesson.MediaItems {
		lessonMedias = append(
			lessonMedias,
			&lessonGRPC.MediaItem{
				Id:        oneMedia.ID,
				LessonId:  oneMedia.LessonID,
				S3Link:    oneMedia.S3Link,
				S3Preview: oneMedia.S3Preview,
				Type:      enumToProtoType(oneMedia.Type),
			},
		)
	}

	return &lessonGRPC.GetOneLessonResponse{
		Id:         lesson.ID,
		BoardId:    lesson.BoardID,
		MeetLink:   lesson.MeetLink,
		StartTime:  timestamppb.New(lesson.StartTime),
		MediaItems: lessonMedias,
		Duration:   lesson.Duration,
		TutorId:    lesson.TutorID,
		UserIds:    lesson.UserIDs,
		Status:     enumToProtoStatus(lesson.Status),
	}
}
