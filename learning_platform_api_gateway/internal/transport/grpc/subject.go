package grpc

import (
	subjectGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/subject"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SubjectClient struct {
	client subjectGRPC.SubjectClient
}

func NewSubjectGrpcConnection(subjectGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		subjectGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewSubjectClient(connection *grpc.ClientConn) *SubjectClient {
	return &SubjectClient{
		client: subjectGRPC.NewSubjectClient(connection),
	}
}
