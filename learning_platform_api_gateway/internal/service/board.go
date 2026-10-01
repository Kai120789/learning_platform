package service

import "learning-platform/api-gateway/internal/dto/boardDto"

type BoardService struct {
	client BoardClient
}

type BoardClient interface {
	GetBoard(boardID int64) (*boardDto.Board, error)
	CreateBoard(title string, tutorID, tutorBoardsLimit int64) (*boardDto.Board, error)
	GetTutorBoards(tutorID int64) ([]boardDto.Board, error)
	DeleteBoard(boardID, tutorID int64) error
	UpdateBoard(title string, tutorID, boardID int64) (*boardDto.Board, error)
}

func NewBoardService(
	client BoardClient,
) *BoardService {
	return &BoardService{
		client: client,
	}
}

func (b *BoardService) GetBoard(boardID int64) (*boardDto.Board, error) {
	// TODO: Добавить проверку доступа к доске (мб во внутреннем сервисе)
	res, err := b.client.GetBoard(boardID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BoardService) CreateBoard(title string, tutorID int64) (*boardDto.Board, error) {
	// TODO: Убрать хардкод
	tutorBoardLimit := 5

	res, err := b.client.CreateBoard(title, tutorID, int64(tutorBoardLimit))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BoardService) GetTutorBoards(tutorID int64) ([]boardDto.Board, error) {
	res, err := b.client.GetTutorBoards(tutorID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BoardService) DeleteBoard(boardID, tutorID int64) error {
	err := b.client.DeleteBoard(boardID, tutorID)
	if err != nil {
		return err
	}

	return nil
}

func (b *BoardService) UpdateBoard(title string, tutorID, boardID int64) (*boardDto.Board, error) {
	res, err := b.client.UpdateBoard(title, tutorID, boardID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
