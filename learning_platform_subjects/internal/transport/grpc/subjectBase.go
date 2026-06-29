package grpc

import (
	"context"
	subjectGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/subject"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/subjects/internal/models"
	"learning-platform/subjects/internal/models/enum"
)

type SubjectBaseService interface {
	GetOneSubject(subjectID int64) (*models.Subject, error)
	GetAllSubjects() ([]models.Subject, error)
}

func (s *SubjectGRPCServer) GetOneSubject(
	ctx context.Context,
	in *subjectGRPC.GetOneSubjectRequest,
) (*subjectGRPC.GetOneSubjectResponse, error) {
	res, err := s.SubjectBaseService.GetOneSubject(in.GetId())
	if err != nil {
		s.logger.Error(
			"failed get one subject",
			zap.Int64("subjectID", in.GetId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed get one subject")
	}

	return &subjectGRPC.GetOneSubjectResponse{
		Id:    res.ID,
		Code:  res.Code,
		Title: res.Title,
		Type:  enumToProtoType(res.Type),
	}, nil
}

func (s *SubjectGRPCServer) GetAllSubjects(
	ctx context.Context,
	in *subjectGRPC.GetAllSubjectsRequest,
) (*subjectGRPC.GetAllSubjectsResponse, error) {
	res, err := s.SubjectBaseService.GetAllSubjects()
	if err != nil {
		s.logger.Error(
			"failed get all subjects",
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed get all subjects")
	}

	var resSubjects []*subjectGRPC.GetOneSubjectResponse
	for _, oneSubject := range res {
		resSubjects = append(resSubjects, &subjectGRPC.GetOneSubjectResponse{
			Id:    oneSubject.ID,
			Code:  oneSubject.Code,
			Title: oneSubject.Title,
			Type:  enumToProtoType(oneSubject.Type),
		})
	}

	return &subjectGRPC.GetAllSubjectsResponse{
		Subjects: resSubjects,
	}, nil
}

func enumToProtoType(subjectType enum.Type) subjectGRPC.SubjectType {
	switch subjectType {
	case enum.TypeEGE:
		return subjectGRPC.SubjectType_EGE
	case enum.TypeOGE:
		return subjectGRPC.SubjectType_OGE
	case enum.TypeImprove:
		return subjectGRPC.SubjectType_IMPROVE
	default:
		return subjectGRPC.SubjectType_MEDIA_TYPE_UNSPECIFIED
	}
}
