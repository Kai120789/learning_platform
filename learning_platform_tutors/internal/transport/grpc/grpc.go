package grpc

import (
	"fmt"
	tutorGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/tutor"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	goGRPC "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"learning-platform/tutors/internal/config"
	"net"
)

type GRPCServer struct {
	server *goGRPC.Server
	config *config.Config
	logger *zap.Logger
}

func New(
	config *config.Config,
	logger *zap.Logger,
	base TutorBaseService,
	review TutorReviewService,
	offer TutorOfferService,
	student TutorStudentService,
	subject TutorSubjectService,
) *GRPCServer {
	gRPCServer := goGRPC.NewServer(goGRPC.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))

	tutorGRPC.RegisterTutorServer(gRPCServer, NewTutorGRPCServer(logger, base, review, offer, student, subject))

	reflection.Register(gRPCServer)

	return &GRPCServer{
		server: gRPCServer,
		logger: logger,
		config: config,
	}
}

func (g *GRPCServer) Run() error {
	const op = "grpcapp.Run"

	listener, err := net.Listen("tcp", g.config.GRPCServerAddress)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := g.server.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
