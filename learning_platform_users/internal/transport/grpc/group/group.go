package group

import (
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
	"go.uber.org/zap"
)

type GroupGRPCServer struct {
	userGRPC.UnimplementedGroupServer
	GroupBaseService GroupBaseService
	GroupUserService GroupUserService
	logger           *zap.Logger
}

func NewGroupGRPCServer(
	groupBase GroupBaseService,
	groupUser GroupUserService,
	logger *zap.Logger,
) userGRPC.GroupServer {
	return &GroupGRPCServer{
		GroupBaseService: groupBase,
		GroupUserService: groupUser,
		logger:           logger,
	}
}
