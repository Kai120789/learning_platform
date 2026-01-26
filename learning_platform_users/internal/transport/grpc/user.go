package grpc

import userGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/user"

type UserService interface{}

type UserGRPCServer struct {
	userGRPC.UnimplementedUserServer
	user UserService
}

func NewUserGRPCServer(user UserService) userGRPC.UserServer {
	return &UserGRPCServer{
		user: user,
	}
}
