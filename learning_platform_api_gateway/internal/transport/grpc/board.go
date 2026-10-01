package grpc

import (
	"context"
	boardGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/board"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learning-platform/api-gateway/internal/dto/boardDto"
	"time"
)

type BoardClient struct {
	client boardGRPC.BoardClient
}

func NewBoardGrpcConnection(boardGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		boardGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewBoardClient(connection *grpc.ClientConn) *BoardClient {
	return &BoardClient{
		client: boardGRPC.NewBoardClient(connection),
	}
}

func (b *BoardClient) GetBoard(boardID int64) (*boardDto.Board, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	board, err := b.client.GetBoard(ctx, &boardGRPC.GetBoardRequest{BoardId: boardID})
	if err != nil {
		return nil, err
	}

	return &boardDto.Board{
		ID:      board.GetBoard().GetId(),
		Title:   board.GetBoard().GetTitle(),
		TutorID: board.GetBoard().GetTutorId(),
	}, nil
}

func (b *BoardClient) CreateBoard(title string, tutorID, tutorBoardsLimit int64) (*boardDto.Board, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	board, err := b.client.CreateBoard(ctx, &boardGRPC.CreateBoardRequest{
		Title:           title,
		TutorId:         tutorID,
		TutorBoardLimit: tutorBoardsLimit,
	})
	if err != nil {
		return nil, err
	}

	return &boardDto.Board{
		ID:      board.GetBoard().GetId(),
		Title:   board.GetBoard().GetTitle(),
		TutorID: board.GetBoard().GetTutorId(),
	}, nil
}

func (b *BoardClient) GetTutorBoards(tutorID int64) ([]boardDto.Board, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	boards, err := b.client.GetTutorBoards(ctx, &boardGRPC.GetTutorBoardsRequest{TutorId: tutorID})
	if err != nil {
		return nil, err
	}

	var resBoards []boardDto.Board
	for _, board := range boards.GetBoards() {
		resBoards = append(resBoards, boardDto.Board{
			ID:      board.GetId(),
			Title:   board.GetTitle(),
			TutorID: board.GetTutorId(),
		})
	}

	return resBoards, nil
}

func (b *BoardClient) DeleteBoard(boardID, tutorID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := b.client.DeleteBoard(ctx, &boardGRPC.DeleteBoardRequest{
		BoardId: boardID,
		TutorId: tutorID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *BoardClient) UpdateBoard(title string, tutorID, boardID int64) (*boardDto.Board, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	board, err := b.client.UpdateBoard(ctx, &boardGRPC.UpdateBoardRequest{
		NewTitle: title,
		TutorId:  tutorID,
		BoardId:  boardID,
	})
	if err != nil {
		return nil, err
	}

	return &boardDto.Board{
		ID:      board.GetBoard().GetId(),
		Title:   board.GetBoard().GetTitle(),
		TutorID: board.GetBoard().GetTutorId(),
	}, nil
}
