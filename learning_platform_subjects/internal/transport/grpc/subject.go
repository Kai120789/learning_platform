package grpc

import (
	subjectGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/subject"
	"go.uber.org/zap"
)

type SubjectGRPCServer struct {
	subjectGRPC.UnimplementedSubjectServer
	SubjectBaseService  SubjectBaseService
	UserSubjectsService UserSubjectsService
	logger              *zap.Logger
}

func NewSubjectGRPCServer(
	logger *zap.Logger,
	subject SubjectBaseService,
	userSubject UserSubjectsService,
) subjectGRPC.SubjectServer {
	return &SubjectGRPCServer{
		logger:              logger,
		SubjectBaseService:  subject,
		UserSubjectsService: userSubject,
	}
}
