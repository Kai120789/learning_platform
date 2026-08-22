package grpc

import (
	tutorGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/tutor"
	"go.uber.org/zap"
)

type TutorGRPCServer struct {
	tutorGRPC.UnimplementedTutorServer
	TutorBaseService    TutorBaseService
	TutorReviewService  TutorReviewService
	TutorOfferService   TutorOfferService
	TutorStudentService TutorStudentService
	TutorSubjectService TutorSubjectService
	logger              *zap.Logger
}

func NewTutorGRPCServer(
	logger *zap.Logger,
	base TutorBaseService,
	review TutorReviewService,
	offer TutorOfferService,
	student TutorStudentService,
	subject TutorSubjectService,
) tutorGRPC.TutorServer {
	return &TutorGRPCServer{
		logger:              logger,
		TutorReviewService:  review,
		TutorBaseService:    base,
		TutorOfferService:   offer,
		TutorStudentService: student,
		TutorSubjectService: subject,
	}
}
