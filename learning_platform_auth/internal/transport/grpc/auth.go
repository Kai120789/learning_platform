package grpc

import authGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/auth"

type AuthService interface {
	Login()
	Register()
	RefreshTokens()
	Logout()
	LogoutAll()
	ChangePassword()
	ForceChangePassword()
	ChangeEmail()
	ForceChangeEmail()
}

type AuthGRPCServer struct {
	authGRPC.UnimplementedAuthServer
	auth AuthService
}

func NewAuthGRPCServer(auth AuthService) authGRPC.AuthServer {
	return &AuthGRPCServer{
		auth: auth,
	}
}
