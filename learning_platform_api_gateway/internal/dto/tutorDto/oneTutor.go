package tutorDto

type OneTutor struct {
	TutorInfo TutorShortInfo `json:"tutor_info"`
	Offers    []TutorOffer   `json:"offers"`
	MyReview  *TutorReview   `json:"my_review"`
}

type OneTutorWithInfo struct {
	TutorInfo TutorShortInfoWithSubjects `json:"tutor_info"`
	Offers    []TutorOfferWithInfo       `json:"offers"`
	MyReview  *TutorReview               `json:"my_review"`
}
