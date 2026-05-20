package grpc

type Client struct {
	UserClient  *UserClient
	AuthClient  *AuthClient
	GroupClient *GroupClient
}

func NewClient(
	userGrpcUrl string,
	authGrpcUrl string,
) (*Client, error) {
	userGrpcConnection, err := NewUserGrpcConnection(userGrpcUrl)
	if err != nil {
		return nil, err
	}

	authGrpcConnection, err := NewAuthGrpcConnection(authGrpcUrl)
	if err != nil {
		return nil, err
	}

	groupGrpcConnection, err := NewGroupGrpcConnection(userGrpcUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		UserClient:  NewUserClient(userGrpcConnection),
		AuthClient:  NewAuthClient(authGrpcConnection),
		GroupClient: NewGroupClient(groupGrpcConnection),
	}, nil
}
