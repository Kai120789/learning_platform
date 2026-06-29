package service

type Service struct {
	SubjectService     *SubjectService
	UserSubjectService *UserSubjectService
}

type Storage struct {
	SubjectStorage     SubjectStorage
	UserSubjectStorage UserSubjectStorage
}

func New(storage *Storage) *Service {
	return &Service{
		SubjectService:     NewSubjectService(storage.SubjectStorage),
		UserSubjectService: NewUserSubjectService(storage.UserSubjectStorage),
	}
}
