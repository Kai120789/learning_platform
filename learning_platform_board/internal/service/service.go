package service

type Service struct {
	BoardService     *BoardService
	LiveBoardService *LiveBoardService
}

type Storage struct {
	PostgresStorage PostgresStorage
	MongoStorage    MongoStorage
}

func New(storage *Storage) *Service {
	return &Service{
		BoardService:     NewBoardService(storage.PostgresStorage),
		LiveBoardService: NewLiveBoardService(storage.MongoStorage),
	}
}
