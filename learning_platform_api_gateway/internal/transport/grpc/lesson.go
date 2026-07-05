package grpc

import (
	"context"
	lessonGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/lesson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"learning-platform/api-gateway/internal/dto/enum"
	"learning-platform/api-gateway/internal/dto/lessonDto"
	"time"
)

type LessonClient struct {
	client lessonGRPC.LessonClient
}

func NewLessonGrpcConnection(lessonGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		lessonGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewLessonClient(connection *grpc.ClientConn) *LessonClient {
	return &LessonClient{
		client: lessonGRPC.NewLessonClient(connection),
	}
}

func (l *LessonClient) GetOneLesson(lessonID int64) (*lessonDto.LessonResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lesson, err := l.client.GetOneLesson(ctx, &lessonGRPC.GetOneLessonRequest{Id: lessonID})
	if err != nil {
		return nil, err
	}

	return grpcLessonResponseToDTO(lesson), nil
}

func (l *LessonClient) GetLessonsByUserId(userID int64) ([]lessonDto.LessonResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lessons, err := l.client.GetLessonsByUserId(ctx, &lessonGRPC.GetLessonsByUserIdRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	var resLessons []lessonDto.LessonResponse
	for _, oneLesson := range lessons.GetLessons() {
		resLessons = append(resLessons, *grpcLessonResponseToDTO(oneLesson))
	}

	return resLessons, nil
}

func (l *LessonClient) CreateLesson(lesson lessonDto.CreateLesson) (*lessonDto.LessonResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var medias []*lessonGRPC.CreateMediaItem
	for _, media := range lesson.MediaItems {
		medias = append(medias, &lessonGRPC.CreateMediaItem{
			S3Link:    media.S3Link,
			S3Preview: media.S3Preview,
			Type:      enumToProtoMediaType(media.Type),
		})
	}

	newLesson, err := l.client.CreateLesson(ctx, &lessonGRPC.CreateLessonRequest{
		BoardId:    lesson.BoardID,
		MeetLink:   lesson.MeetLink,
		StartTime:  timestamppb.New(lesson.StartTime),
		Duration:   lesson.Duration,
		TutorId:    lesson.TutorID,
		MediaItems: medias,
		UserIds:    lesson.UserIDs,
	})
	if err != nil {
		return nil, err
	}

	return grpcLessonResponseToDTO(newLesson.GetLesson()), nil
}

func (l *LessonClient) UpdateLesson(lesson lessonDto.UpdateLesson) (*lessonDto.LessonResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var medias []*lessonGRPC.CreateMediaItem
	for _, media := range lesson.MediaItems {
		medias = append(medias, &lessonGRPC.CreateMediaItem{
			S3Link:    media.S3Link,
			S3Preview: media.S3Preview,
			Type:      enumToProtoMediaType(media.Type),
		})
	}

	newLesson, err := l.client.UpdateLesson(ctx, &lessonGRPC.UpdateLessonRequest{
		Id:              lesson.ID,
		BoardId:         lesson.BoardID,
		MeetLink:        lesson.MeetLink,
		StartTime:       timestamppb.New(lesson.StartTime),
		Duration:        lesson.Duration,
		MediaItems:      medias,
		UserIds:         lesson.UserIDs,
		DeletedMediaIds: lesson.DeletedMediaIDs,
		DeletedUserIds:  lesson.DeletedUserIDs,
	})
	if err != nil {
		return nil, err
	}

	return grpcLessonResponseToDTO(newLesson.GetLesson()), nil
}

func (l *LessonClient) UpdateLessonStatus(lessonID int64, status enum.LessonStatus) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := l.client.UpdateLessonStatus(ctx, &lessonGRPC.UpdateLessonStatusRequest{
		Id:     lessonID,
		Status: enumToProtoLessonStatus(status),
	})
	if err != nil {
		return err
	}

	return nil
}

func (l *LessonClient) GetLessonsByTutorId(tutorID int64) ([]lessonDto.LessonResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lessons, err := l.client.GetLessonsByTutorId(ctx, &lessonGRPC.GetLessonsByTutorIdRequest{TutorId: tutorID})
	if err != nil {
		return nil, err
	}

	var resLessons []lessonDto.LessonResponse
	for _, oneLesson := range lessons.GetLessons() {
		resLessons = append(resLessons, *grpcLessonResponseToDTO(oneLesson))
	}

	return resLessons, nil
}

func grpcLessonResponseToDTO(lesson *lessonGRPC.GetOneLessonResponse) *lessonDto.LessonResponse {
	var lessonMedias []lessonDto.MediaItemResponse

	for _, oneMedia := range lesson.GetMediaItems() {
		lessonMedias = append(lessonMedias, lessonDto.MediaItemResponse{
			ID:        oneMedia.GetId(),
			LessonID:  oneMedia.GetLessonId(),
			S3Link:    oneMedia.GetS3Link(),
			S3Preview: oneMedia.GetS3Preview(),
			Type:      protoToEnumMediaType(oneMedia.GetType()),
		})
	}

	return &lessonDto.LessonResponse{
		ID:         lesson.GetId(),
		BoardID:    lesson.BoardId,
		MeetLink:   lesson.MeetLink,
		StartTime:  lesson.GetStartTime().AsTime(),
		Duration:   lesson.GetDuration(),
		TutorID:    lesson.GetTutorId(),
		Status:     protoToEnumLessonStatus(lesson.GetStatus()),
		MediaItems: lessonMedias,
		UserIDs:    lesson.GetUserIds(),
	}
}

func protoToEnumMediaType(mediaType lessonGRPC.MediaType) enum.MediaType {
	switch mediaType {
	case lessonGRPC.MediaType_IMAGE:
		return enum.TypeImage
	case lessonGRPC.MediaType_VIDEO:
		return enum.TypeVideo
	default:
		return ""
	}
}

func enumToProtoMediaType(mediaType enum.MediaType) lessonGRPC.MediaType {
	switch mediaType {
	case enum.TypeImage:
		return lessonGRPC.MediaType_IMAGE
	case enum.TypeVideo:
		return lessonGRPC.MediaType_VIDEO
	default:
		return lessonGRPC.MediaType_MEDIA_TYPE_UNSPECIFIED
	}
}

func protoToEnumLessonStatus(status lessonGRPC.LessonStatus) enum.LessonStatus {
	switch status {
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

func enumToProtoLessonStatus(status enum.LessonStatus) lessonGRPC.LessonStatus {
	switch status {
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
