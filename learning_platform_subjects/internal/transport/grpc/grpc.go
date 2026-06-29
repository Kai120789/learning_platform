package grpc

import (
	"fmt"
	subjectGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/subject"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	goGRPC "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"learning-platform/subjects/internal/config"
	"net"
)

type GRPCServer struct {
	server *goGRPC.Server
	config *config.Config
	logger *zap.Logger
}

func New(
	logger *zap.Logger,
	config *config.Config,
	subject SubjectBaseService,
	userSubjects UserSubjectsService,
) *GRPCServer {
	gRPCServer := goGRPC.NewServer(goGRPC.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))

	subjectGRPC.RegisterSubjectServer(
		gRPCServer,
		NewSubjectGRPCServer(
			logger,
			subject,
			userSubjects,
		),
	)

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
