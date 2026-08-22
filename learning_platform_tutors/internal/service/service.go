package service

type Service struct {
	TutorService        *TutorService
	TutorReviewService  *TutorReviewService
	TutorOfferService   *TutorOfferService
	TutorStudentService *TutorStudentService
	TutorSubjectService *TutorSubjectService
}

type Storage struct {
	TutorReviewStorage  TutorReviewStorage
	TutorOfferStorage   TutorOfferStorage
	TutorStudentStorage TutorStudentStorage
	TutorSubjectStorage TutorSubjectStorage
}

func New(storage *Storage) *Service {
	reviewService := NewTutorReviewService(storage.TutorReviewStorage)
	offerService := NewTutorOfferService(storage.TutorOfferStorage)
	studentService := NewTutorStudentService(storage.TutorStudentStorage)
	subjectService := NewTutorSubjectService(storage.TutorSubjectStorage)

	return &Service{
		TutorService: NewTutorService(
			reviewService,
			offerService,
			studentService,
			subjectService,
		),
		TutorReviewService:  reviewService,
		TutorOfferService:   offerService,
		TutorStudentService: studentService,
		TutorSubjectService: subjectService,
	}
}
