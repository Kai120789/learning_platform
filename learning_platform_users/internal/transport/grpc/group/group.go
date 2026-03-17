package group

import (
	userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"
)

type GroupGRPCServer struct {
	userGRPC.UnimplementedGroupServer
	GroupBaseService GroupBaseService
	GroupUserService GroupUserService
}

func NewGroupGRPCServer(
	groupBase GroupBaseService,
	groupUser GroupUserService,
) userGRPC.GroupServer {
	return &GroupGRPCServer{
		GroupBaseService: groupBase,
		GroupUserService: groupUser,
	}
}
