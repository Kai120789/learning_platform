package grpc

import (
	"context"
	subjectGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/subject"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/subjects/internal/models"
)

type UserSubjectsService interface {
	GetUserSubjects(userID int64) ([]models.Subject, error)
	SetUserSubjects(userID int64, subjectIDs []int64) ([]models.Subject, error)
	UpdateUserSubjects(userID int64, subjectIDs []int64, deletedSubjectIDs []int64) ([]models.Subject, error)
}

func (s *SubjectGRPCServer) GetUserSubjects(
	ctx context.Context,
	in *subjectGRPC.GetUserSubjectsRequest,
) (*subjectGRPC.GetUserSubjectsResponse, error) {
	res, err := s.UserSubjectsService.GetUserSubjects(in.GetUserId())
	if err != nil {
		s.logger.Error(
			"failed get user subjects",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed get users subjects")
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

	return &subjectGRPC.GetUserSubjectsResponse{
		Subjects: resSubjects,
	}, nil
}

func (s *SubjectGRPCServer) SetUserSubjects(
	ctx context.Context,
	in *subjectGRPC.SetUserSubjectsRequest,
) (*subjectGRPC.SetUserSubjectsResponse, error) {
	res, err := s.UserSubjectsService.SetUserSubjects(in.GetUserId(), in.GetSubjectIds())
	if err != nil {
		s.logger.Error(
			"failed set user subjects",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed set users subjects")
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

	return &subjectGRPC.SetUserSubjectsResponse{
		Subjects: resSubjects,
	}, nil
}

func (s *SubjectGRPCServer) UpdateUserSubjects(
	ctx context.Context,
	in *subjectGRPC.UpdateUserSubjectsRequest,
) (*subjectGRPC.UpdateUserSubjectsResponse, error) {
	res, err := s.UserSubjectsService.UpdateUserSubjects(
		in.GetUserId(),
		in.GetSubjectIds(),
		in.GetDeletedSubjectIds(),
	)
	if err != nil {
		s.logger.Error(
			"failed set user subjects",
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed set users subjects")
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

	return &subjectGRPC.UpdateUserSubjectsResponse{
		Subjects: resSubjects,
	}, nil
}
