package grpc

import (
	groupGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
)

type GroupGRPCServer struct {
	groupGRPC.UnimplementedGroupServer
	service *GroupService
	logger  *zap.Logger
}

type GroupService struct {
	GroupBaseService GroupBaseService
	GroupUserService GroupUserService
}

func NewGroupGRPCServer(
	service *GroupService,
	logger *zap.Logger,
) groupGRPC.GroupServer {
	return &GroupGRPCServer{
		service: service,
		logger:  logger,
	}
}
