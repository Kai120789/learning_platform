package grpc

import (
	"fmt"
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	goGRPC "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"learning-platform/users/internal/config"
	groupGrpcServer "learning-platform/users/internal/transport/grpc/group"
	userGrpcServer "learning-platform/users/internal/transport/grpc/user"
	"net"
)

type GRPCServer struct {
	server *goGRPC.Server
	logger *zap.Logger
	config *config.Config
}

func New(
	logger *zap.Logger,
	config *config.Config,
	user *userGrpcServer.UserGRPCServer,
	group *groupGrpcServer.GroupGRPCServer,
) *GRPCServer {
	gRPCServer := goGRPC.NewServer(goGRPC.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))

	userGRPC.RegisterUserServer(
		gRPCServer,
		userGrpcServer.NewUserGRPCServer(
			user.UserBaseService,
			user.UserSettingsService,
			user.UserInfoService,
			logger,
		),
	)

	userGRPC.RegisterGroupServer(
		gRPCServer,
		groupGrpcServer.NewGroupGRPCServer(
			group.GroupBaseService,
			group.GroupUserService,
			logger,
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
