package service

import postgresModels "learning-platform/board/internal/models/postgres"

type BoardService struct {
	postgres PostgresStorage
}

type PostgresStorage interface {
	GetBoard(boardID int64) (*postgresModels.Board, error)
	CreateBoard(title string, tutorID, tutorBoardsLimit int64) (*postgresModels.Board, error)
	GetTutorBoards(tutorID int64) ([]postgresModels.Board, error)
	DeleteBoard(boardID, tutorID int64) error
	UpdateBoard(title string, tutorID, boardID int64) (*postgresModels.Board, error)
}

func NewBoardService(postgres PostgresStorage) *BoardService {
	return &BoardService{
		postgres: postgres,
	}
}

func (b *BoardService) GetBoard(boardID int64) (*postgresModels.Board, error) {
	res, err := b.postgres.GetBoard(boardID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BoardService) CreateBoard(title string, tutorID, tutorBoardsLimit int64) (*postgresModels.Board, error) {
	res, err := b.postgres.CreateBoard(title, tutorID, tutorBoardsLimit)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BoardService) GetTutorBoards(tutorID int64) ([]postgresModels.Board, error) {
	res, err := b.postgres.GetTutorBoards(tutorID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BoardService) DeleteBoard(boardID, tutorID int64) error {
	err := b.postgres.DeleteBoard(boardID, tutorID)
	if err != nil {
		return err
	}

	return nil
}

func (b *BoardService) UpdateBoard(title string, tutorID, boardID int64) (*postgresModels.Board, error) {
	res, err := b.postgres.UpdateBoard(title, tutorID, boardID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
