package grpc

import (
	"fmt"
	lessonGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/lesson"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	goGRPC "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"learning-platform/lessons/internal/config"
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
	lesson LessonService,
) *GRPCServer {
	gRPCServer := goGRPC.NewServer(goGRPC.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))

	lessonGRPC.RegisterLessonServer(
		gRPCServer,
		NewLessonGRPCServer(
			logger,
			lesson,
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
