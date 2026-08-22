package grpc

import (
	"context"
	tutorGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/tutor"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TutorSubjectService interface {
	AddTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error)
	UpdateTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error)
}

func (t *TutorGRPCServer) AddTutorSubjects(
	ctx context.Context,
	in *tutorGRPC.AddTutorSubjectsRequest,
) (*tutorGRPC.AddTutorSubjectsResponse, error) {
	subjectIDs, err := t.TutorSubjectService.AddTutorSubjects(in.GetTutorId(), in.GetSubjectIds())
	if err != nil {
		t.logger.Error(
			"failed to add subjects",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Int64s("subjectIDs", in.GetSubjectIds()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to add subjects")
	}

	return &tutorGRPC.AddTutorSubjectsResponse{
		SubjectIds: subjectIDs,
	}, nil
}

func (t *TutorGRPCServer) UpdateTutorSubjects(
	ctx context.Context,
	in *tutorGRPC.UpdateTutorSubjectsRequest,
) (*tutorGRPC.UpdateTutorSubjectsResponse, error) {
	subjectIDs, err := t.TutorSubjectService.UpdateTutorSubjects(
		in.GetTutorId(),
		in.GetSubjectIds(),
	)
	if err != nil {
		t.logger.Error(
			"failed to add subjects",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Int64s("subjectIDs", in.GetSubjectIds()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to add subjects")
	}

	return &tutorGRPC.UpdateTutorSubjectsResponse{
		SubjectIds: subjectIDs,
	}, nil
}
