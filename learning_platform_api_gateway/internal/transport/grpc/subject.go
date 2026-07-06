package grpc

import (
	"context"
	subjectGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/subject"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learning-platform/api-gateway/internal/dto/enum"
	"learning-platform/api-gateway/internal/dto/subjectDto"
	"time"
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

func (s *SubjectClient) GetOneSubject(subjectID int64) (*subjectDto.Subject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetOneSubject(ctx, &subjectGRPC.GetOneSubjectRequest{Id: subjectID})
	if err != nil {
		return nil, err
	}

	return mapGrpcSubjectToDTO(res), nil
}

func (s *SubjectClient) GetAllSubjects() ([]subjectDto.Subject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetAllSubjects(ctx, &subjectGRPC.GetAllSubjectsRequest{})
	if err != nil {
		return nil, err
	}

	var resSubjects []subjectDto.Subject
	for _, oneSubject := range res.GetSubjects() {
		resSubjects = append(resSubjects, *mapGrpcSubjectToDTO(oneSubject))
	}

	return resSubjects, nil
}

func (s *SubjectClient) GetUserSubjects(userID int64) ([]subjectDto.Subject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetUserSubjects(ctx, &subjectGRPC.GetUserSubjectsRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	var resSubjects []subjectDto.Subject
	for _, oneSubject := range res.GetSubjects() {
		resSubjects = append(resSubjects, *mapGrpcSubjectToDTO(oneSubject))
	}

	return resSubjects, nil
}

func (s *SubjectClient) SetUserSubjects(userID int64, subjectIDs []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.SetUserSubjects(ctx, &subjectGRPC.SetUserSubjectsRequest{
		UserId:     userID,
		SubjectIds: subjectIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *SubjectClient) UpdateUserSubjects(userID int64, subjectIDs, deletedSubjectIDs []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.UpdateUserSubjects(ctx, &subjectGRPC.UpdateUserSubjectsRequest{
		UserId:            userID,
		SubjectIds:        subjectIDs,
		DeletedSubjectIds: deletedSubjectIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

func mapGrpcSubjectToDTO(subject *subjectGRPC.GetOneSubjectResponse) *subjectDto.Subject {
	return &subjectDto.Subject{
		ID:    subject.GetId(),
		Code:  subject.GetCode(),
		Title: subject.GetTitle(),
		Type:  protoToEnumType(subject.GetType()),
	}
}

func protoToEnumType(subjectType subjectGRPC.SubjectType) enum.SubjectType {
	switch subjectType {
	case subjectGRPC.SubjectType_EGE:
		return enum.TypeEGE
	case subjectGRPC.SubjectType_OGE:
		return enum.TypeOGE
	case subjectGRPC.SubjectType_IMPROVE:
		return enum.TypeImprove
	default:
		return ""
	}
}
