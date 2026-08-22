package grpc

import (
	"context"
	tutorGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/tutor"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/tutors/internal/dto"
	"learning-platform/tutors/internal/models"
	"learning-platform/tutors/internal/utils"
)

type TutorStudentService interface {
	AddStudent(tutorID, studentID int64) error
	DeleteOneStudent(tutorID, studentID int64) error
	DeleteStudents(tutorID int64, studentIDs []int64) error
	GetTutorStudents(getTutorStudents dto.GetTutorStudents) ([]models.TutorStudent, int64, error)
}

func (t *TutorGRPCServer) AddStudent(
	ctx context.Context,
	in *tutorGRPC.AddStudentRequest,
) (*tutorGRPC.AddStudentResponse, error) {
	err := t.TutorStudentService.AddStudent(in.GetTutorId(), in.GetStudentId())
	if err != nil {
		t.logger.Error(
			"failed to add student",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Int64("studentID", in.GetStudentId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to add student")
	}

	return &tutorGRPC.AddStudentResponse{}, nil
}

func (t *TutorGRPCServer) DeleteOneStudent(
	ctx context.Context,
	in *tutorGRPC.DeleteOneStudentRequest,
) (*tutorGRPC.DeleteOneStudentResponse, error) {
	err := t.TutorStudentService.DeleteOneStudent(in.GetTutorId(), in.GetStudentId())
	if err != nil {
		t.logger.Error(
			"failed to delete one student",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Int64("studentID", in.GetStudentId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete one student")
	}

	return &tutorGRPC.DeleteOneStudentResponse{}, nil
}

func (t *TutorGRPCServer) DeleteStudents(
	ctx context.Context,
	in *tutorGRPC.DeleteStudentsRequest,
) (*tutorGRPC.DeleteStudentsResponse, error) {
	err := t.TutorStudentService.DeleteStudents(in.GetTutorId(), in.GetStudentIds())
	if err != nil {
		t.logger.Error(
			"failed to delete students",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Int64s("studentIDs", in.GetStudentIds()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete students")
	}

	return &tutorGRPC.DeleteStudentsResponse{}, nil
}

func (t *TutorGRPCServer) GetTutorStudents(
	ctx context.Context,
	in *tutorGRPC.GetTutorStudentsRequest,
) (*tutorGRPC.GetTutorStudentsResponse, error) {
	request := dto.GetTutorStudents{
		TutorID:              in.GetTutorId(),
		InteractedWithinDays: in.InteractedWithinDays,
		Page:                 in.GetPage(),
		Limit:                in.GetLimit(),
	}

	students, count, err := t.TutorStudentService.GetTutorStudents(request)
	if err != nil {
		t.logger.Error(
			"failed to get tutor students",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get tutor students")
	}

	var resStudents []*tutorGRPC.TutorStudent
	for _, oneStudent := range students {
		resStudents = append(resStudents, &tutorGRPC.TutorStudent{
			StudentId:        oneStudent.StudentID,
			LastInteractedAt: utils.DBTimeToOptional(oneStudent.LastInteractedAt),
		})
	}

	return &tutorGRPC.GetTutorStudentsResponse{
		Students: resStudents,
		Count:    count,
	}, nil
}
