package grpc

import (
	"context"
	boardGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/board"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	postgresModels "learning-platform/board/internal/models/postgres"
)

type BoardGRPCServer struct {
	boardGRPC.UnimplementedBoardServer
	service Service
	logger  *zap.Logger
}

func NewBoardGRPCServer(
	logger *zap.Logger,
	service Service,
) boardGRPC.BoardServer {
	return &BoardGRPCServer{
		logger:  logger,
		service: service,
	}
}

type Service interface {
	GetBoard(boardID int64) (*postgresModels.Board, error)
	CreateBoard(title string, tutorID, tutorBoardsLimit int64) (*postgresModels.Board, error)
	GetTutorBoards(tutorID int64) ([]postgresModels.Board, error)
	DeleteBoard(boardID, tutorID int64) error
	UpdateBoard(title string, tutorID, boardID int64) (*postgresModels.Board, error)
}

func (b *BoardGRPCServer) GetBoard(
	ctx context.Context,
	in *boardGRPC.GetBoardRequest,
) (*boardGRPC.GetBoardResponse, error) {
	board, err := b.service.GetBoard(in.GetBoardId())
	if err != nil {
		b.logger.Error(
			"failed to get board",
			zap.Int64("boardID", in.GetBoardId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get board")
	}

	return &boardGRPC.GetBoardResponse{
		Board: &boardGRPC.OneBoard{
			Id:      board.ID,
			Title:   board.Title,
			TutorId: board.TutorID,
		},
	}, nil
}

func (b *BoardGRPCServer) CreateBoard(
	ctx context.Context,
	in *boardGRPC.CreateBoardRequest,
) (*boardGRPC.CreateBoardResponse, error) {
	board, err := b.service.CreateBoard(in.GetTitle(), in.GetTutorId(), in.GetTutorBoardLimit())
	if err != nil {
		b.logger.Error(
			"failed to create board",
			zap.String("title", in.GetTitle()),
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to create board")
	}

	return &boardGRPC.CreateBoardResponse{
		Board: &boardGRPC.OneBoard{
			Id:      board.ID,
			Title:   board.Title,
			TutorId: board.TutorID,
		},
	}, nil
}

func (b *BoardGRPCServer) GetTutorBoards(
	ctx context.Context,
	in *boardGRPC.GetTutorBoardsRequest,
) (*boardGRPC.GetTutorBoardsResponse, error) {
	boards, err := b.service.GetTutorBoards(in.GetTutorId())
	if err != nil {
		b.logger.Error(
			"failed to get tutor boards",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get tutor boards")
	}

	var resBoards []*boardGRPC.OneBoard
	for _, oneBoard := range boards {
		resBoards = append(resBoards, &boardGRPC.OneBoard{
			Id:      oneBoard.ID,
			Title:   oneBoard.Title,
			TutorId: oneBoard.TutorID,
		})
	}

	return &boardGRPC.GetTutorBoardsResponse{
		Boards: resBoards,
	}, nil
}

func (b *BoardGRPCServer) DeleteBoard(
	ctx context.Context,
	in *boardGRPC.DeleteBoardRequest,
) (*boardGRPC.DeleteBoardResponse, error) {
	err := b.service.DeleteBoard(in.GetBoardId(), in.GetTutorId())
	if err != nil {
		b.logger.Error(
			"failed to delete board",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Int64("boardID", in.GetBoardId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete board")
	}

	return &boardGRPC.DeleteBoardResponse{}, nil
}

func (b *BoardGRPCServer) UpdateBoard(
	ctx context.Context,
	in *boardGRPC.UpdateBoardRequest,
) (*boardGRPC.UpdateBoardResponse, error) {
	board, err := b.service.UpdateBoard(in.GetNewTitle(), in.GetTutorId(), in.GetBoardId())
	if err != nil {
		b.logger.Error(
			"failed to update board",
			zap.String("title", in.GetNewTitle()),
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Int64("boardID", in.GetBoardId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update board")
	}

	return &boardGRPC.UpdateBoardResponse{
		Board: &boardGRPC.OneBoard{
			Id:      board.ID,
			Title:   board.Title,
			TutorId: board.TutorID,
		},
	}, nil
}
