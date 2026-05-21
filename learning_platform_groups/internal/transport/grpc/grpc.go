package grpc

import (
	"fmt"
	groupGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	goGRPC "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"learning-platform/groups/internal/config"
	"net"
)

type GRPCServer struct {
	server *goGRPC.Server
	config *config.Config
	logger *zap.Logger
}

func New(config *config.Config, logger *zap.Logger, service *GroupService) *GRPCServer {
	gRPCServer := goGRPC.NewServer(goGRPC.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))

	groupGRPC.RegisterGroupServer(gRPCServer, NewGroupGRPCServer(service, logger))

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
