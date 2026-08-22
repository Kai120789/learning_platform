package service

import (
	"learning-platform/api-gateway/internal/dto/subjectDto"
	"learning-platform/api-gateway/internal/dto/tutorDto"
	"learning-platform/api-gateway/internal/dto/userDto"
)

type TutorService struct {
	client  TutorClient
	subject TutorSubjectService
	user    TutorUserService
}

type TutorClient interface {
	AddStudent(tutorID, studentID int64) error
	DeleteOneStudent(tutorID, studentID int64) error
	DeleteStudents(tutorID int64, studentIDs []int64) error
	GetTutorStudents(getTutorStudents tutorDto.GetTutorStudents) ([]tutorDto.TutorStudent, int64, error)
	GetOneTutor(tutorID, userID int64) (*tutorDto.OneTutor, error)
	GetTutors(getTutors tutorDto.GetTutors) ([]tutorDto.TutorShortInfo, int64, error)
	GetTutorReviews(getTutorReviews tutorDto.GetTutorReviews) ([]tutorDto.TutorReview, int64, error)
	AddTutorReview(newReview tutorDto.NewTutorReview) (*tutorDto.TutorReview, error)
	UpdateTutorReview(updReview tutorDto.UpdateTutorReview) (*tutorDto.TutorReview, error)
	DeleteTutorReview(reviewID, authorID int64) error
	GetTutorOffers(tutorID int64) ([]tutorDto.TutorOffer, error)
	AddTutorOffer(newOffer tutorDto.NewTutorOffer) (*tutorDto.TutorOffer, error)
	UpdateTutorOffer(updOffer tutorDto.UpdateTutorOffer) (*tutorDto.TutorOffer, error)
	DeleteOneTutorOffer(offerID, tutorID int64) error
	DeleteTutorOffers(offerIDs []int64, tutorID int64) error
	AddTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error)
	UpdateTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error)
}

type TutorSubjectService interface {
	GetAllSubjects() ([]subjectDto.Subject, error)
	GetOneSubject(subjectID int64) (*subjectDto.Subject, error)
}

type TutorUserService interface {
	GetUsersShortInfo(userIDs []int64) ([]userDto.UserShortInfo, error)
}

func NewTutorService(
	client TutorClient,
	subject TutorSubjectService,
	user TutorUserService,
) *TutorService {
	return &TutorService{
		client:  client,
		subject: subject,
		user:    user,
	}
}

func (t *TutorService) AddStudent(tutorID, studentID int64) error {
	err := t.client.AddStudent(tutorID, studentID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorService) DeleteOneStudent(tutorID, studentID int64) error {
	err := t.client.DeleteOneStudent(tutorID, studentID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorService) DeleteStudents(tutorID int64, studentIDs []int64) error {
	err := t.client.DeleteStudents(tutorID, studentIDs)
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorService) GetTutorStudents(getTutorStudents tutorDto.GetTutorStudents) ([]tutorDto.TutorStudentWithInfo, int64, error) {
	tutorStudents, count, err := t.client.GetTutorStudents(getTutorStudents)
	if err != nil {
		return nil, 0, err
	}

	var studentIDs []int64
	for _, oneStudent := range tutorStudents {
		studentIDs = append(studentIDs, oneStudent.StudentID)
	}

	students, err := t.user.GetUsersShortInfo(studentIDs)
	if err != nil {
		return nil, 0, err
	}

	studentsByID := make(map[int64]userDto.UserShortInfo)
	for _, oneStudent := range students {
		studentsByID[oneStudent.ID] = oneStudent
	}

	var resTutorStudents []tutorDto.TutorStudentWithInfo
	for _, oneStudent := range tutorStudents {
		resTutorStudents = append(resTutorStudents, tutorDto.TutorStudentWithInfo{
			Student:          studentsByID[oneStudent.StudentID],
			LastInteractedAt: oneStudent.LastInteractedAt,
		})
	}

	return resTutorStudents, count, nil
}

func (t *TutorService) GetOneTutor(tutorID, userID int64) (*tutorDto.OneTutorWithInfo, error) {
	tutor, err := t.client.GetOneTutor(tutorID, userID)
	if err != nil {
		return nil, err
	}

	subjects, err := t.subject.GetAllSubjects()
	if err != nil {
		return nil, err
	}

	subjectsByID := make(map[int64]subjectDto.Subject)
	for _, oneSubject := range subjects {
		subjectsByID[oneSubject.ID] = oneSubject
	}

	resTutor, err := t.mapTutorsWithInfoDTO([]tutorDto.TutorShortInfo{tutor.TutorInfo}, subjectsByID)
	if err != nil {
		return nil, err
	}

	var offersWithInfo []tutorDto.TutorOfferWithInfo
	for _, oneOffer := range tutor.Offers {
		offersWithInfo = append(offersWithInfo, tutorDto.TutorOfferWithInfo{
			ID:              oneOffer.ID,
			TutorID:         oneOffer.TutorID,
			Subject:         subjectsByID[oneOffer.SubjectID],
			Title:           oneOffer.Title,
			Description:     oneOffer.Description,
			Price:           oneOffer.Price,
			DurationMinutes: oneOffer.DurationMinutes,
		})
	}

	return &tutorDto.OneTutorWithInfo{
		TutorInfo: resTutor[0],
		MyReview:  tutor.MyReview,
		Offers:    offersWithInfo,
	}, nil
}

func (t *TutorService) GetTutors(getTutors tutorDto.GetTutors) ([]tutorDto.TutorShortInfoWithSubjects, int64, error) {
	res, count, err := t.client.GetTutors(getTutors)
	if err != nil {
		return nil, 0, err
	}

	subjects, err := t.subject.GetAllSubjects()
	if err != nil {
		return nil, 0, err
	}

	subjectsByID := make(map[int64]subjectDto.Subject)
	for _, oneSubject := range subjects {
		subjectsByID[oneSubject.ID] = oneSubject
	}

	resTutors, err := t.mapTutorsWithInfoDTO(res, subjectsByID)
	if err != nil {
		return nil, 0, err
	}

	return resTutors, count, nil
}

func (t *TutorService) GetTutorReviews(getTutorReviews tutorDto.GetTutorReviews) ([]tutorDto.TutorReviewWithInfo, int64, error) {
	reviews, count, err := t.client.GetTutorReviews(getTutorReviews)
	if err != nil {
		return nil, 0, err
	}

	var reviewsAuthorIDs, tutorIDs []int64
	for _, oneReview := range reviews {
		reviewsAuthorIDs = append(reviewsAuthorIDs, oneReview.AuthorID)
		tutorIDs = append(tutorIDs, oneReview.TutorID)
	}

	reviewsAuthors, err := t.user.GetUsersShortInfo(reviewsAuthorIDs)
	if err != nil {
		return nil, 0, err
	}

	authorsByID := make(map[int64]userDto.UserShortInfo)
	for _, oneAuthor := range reviewsAuthors {
		authorsByID[oneAuthor.ID] = oneAuthor
	}

	subjects, err := t.subject.GetAllSubjects()
	if err != nil {
		return nil, 0, err
	}

	subjectsByID := make(map[int64]subjectDto.Subject)
	for _, oneSubject := range subjects {
		subjectsByID[oneSubject.ID] = oneSubject
	}

	tutorsInfo, err := t.user.GetUsersShortInfo(tutorIDs)
	if err != nil {
		return nil, 0, err
	}

	tutorsByID := make(map[int64]userDto.UserShortInfo)
	for _, oneTutor := range tutorsInfo {
		tutorsByID[oneTutor.ID] = oneTutor
	}

	var resReviews []tutorDto.TutorReviewWithInfo
	for _, oneReview := range reviews {
		resReviews = append(resReviews, tutorDto.TutorReviewWithInfo{
			ID:        oneReview.ID,
			Tutor:     tutorsByID[oneReview.TutorID],
			Author:    authorsByID[oneReview.AuthorID],
			Subject:   subjectsByID[oneReview.SubjectID],
			Text:      oneReview.Text,
			Rating:    oneReview.Rating,
			CreatedAt: oneReview.CreatedAt,
			UpdatedAt: oneReview.UpdatedAt,
		})
	}

	return resReviews, count, nil
}

func (t *TutorService) AddTutorReview(newReview tutorDto.NewTutorReview) (*tutorDto.TutorReviewWithInfo, error) {
	review, err := t.client.AddTutorReview(newReview)
	if err != nil {
		return nil, err
	}

	authorInfo, err := t.user.GetUsersShortInfo([]int64{review.AuthorID})
	if err != nil || len(authorInfo) == 0 {
		return nil, err
	}

	tutorInfo, err := t.user.GetUsersShortInfo([]int64{review.TutorID})
	if err != nil || len(tutorInfo) == 0 {
		return nil, err
	}

	subject, err := t.subject.GetOneSubject(review.SubjectID)
	if err != nil {
		return nil, err
	}

	return &tutorDto.TutorReviewWithInfo{
		ID:        review.ID,
		Tutor:     tutorInfo[0],
		Author:    authorInfo[0],
		Subject:   *subject,
		Text:      review.Text,
		Rating:    review.Rating,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}, nil
}

func (t *TutorService) UpdateTutorReview(updReview tutorDto.UpdateTutorReview) (*tutorDto.TutorReviewWithInfo, error) {
	review, err := t.client.UpdateTutorReview(updReview)
	if err != nil {
		return nil, err
	}
	authorInfo, err := t.user.GetUsersShortInfo([]int64{review.AuthorID})
	if err != nil || len(authorInfo) == 0 {
		return nil, err
	}

	tutorInfo, err := t.user.GetUsersShortInfo([]int64{review.TutorID})
	if err != nil || len(tutorInfo) == 0 {
		return nil, err
	}

	subject, err := t.subject.GetOneSubject(review.SubjectID)
	if err != nil {
		return nil, err
	}

	return &tutorDto.TutorReviewWithInfo{
		ID:        review.ID,
		Tutor:     tutorInfo[0],
		Author:    authorInfo[0],
		Subject:   *subject,
		Text:      review.Text,
		Rating:    review.Rating,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}, nil
}

func (t *TutorService) DeleteTutorReview(reviewID, authorID int64) error {
	err := t.client.DeleteTutorReview(reviewID, authorID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorService) GetTutorOffers(tutorID int64) ([]tutorDto.TutorOfferWithInfo, error) {
	offers, err := t.client.GetTutorOffers(tutorID)
	if err != nil {
		return nil, err
	}

	subjects, err := t.subject.GetAllSubjects()
	if err != nil {
		return nil, err
	}

	subjectsByID := make(map[int64]subjectDto.Subject)
	for _, oneSubject := range subjects {
		subjectsByID[oneSubject.ID] = oneSubject
	}

	var offersWithInfo []tutorDto.TutorOfferWithInfo
	for _, oneOffer := range offers {
		offersWithInfo = append(offersWithInfo, tutorDto.TutorOfferWithInfo{
			ID:              oneOffer.ID,
			TutorID:         oneOffer.TutorID,
			Subject:         subjectsByID[oneOffer.SubjectID],
			Title:           oneOffer.Title,
			Description:     oneOffer.Description,
			Price:           oneOffer.Price,
			DurationMinutes: oneOffer.DurationMinutes,
		})
	}

	return offersWithInfo, nil
}

func (t *TutorService) AddTutorOffer(newOffer tutorDto.NewTutorOffer) (*tutorDto.TutorOfferWithInfo, error) {
	offer, err := t.client.AddTutorOffer(newOffer)
	if err != nil {
		return nil, err
	}

	subject, err := t.subject.GetOneSubject(offer.SubjectID)
	if err != nil {
		return nil, err
	}

	return &tutorDto.TutorOfferWithInfo{
		ID:              offer.ID,
		TutorID:         offer.TutorID,
		Subject:         *subject,
		Title:           offer.Title,
		Description:     offer.Description,
		Price:           offer.Price,
		DurationMinutes: offer.DurationMinutes,
	}, nil
}

func (t *TutorService) UpdateTutorOffer(updOffer tutorDto.UpdateTutorOffer) (*tutorDto.TutorOfferWithInfo, error) {
	offer, err := t.client.UpdateTutorOffer(updOffer)
	if err != nil {
		return nil, err
	}

	subject, err := t.subject.GetOneSubject(offer.SubjectID)
	if err != nil {
		return nil, err
	}

	return &tutorDto.TutorOfferWithInfo{
		ID:              offer.ID,
		TutorID:         offer.TutorID,
		Subject:         *subject,
		Title:           offer.Title,
		Description:     offer.Description,
		Price:           offer.Price,
		DurationMinutes: offer.DurationMinutes,
	}, nil
}

func (t *TutorService) DeleteOneTutorOffer(offerID, tutorID int64) error {
	err := t.client.DeleteOneTutorOffer(offerID, tutorID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorService) DeleteTutorOffers(offerIDs []int64, tutorID int64) error {
	err := t.client.DeleteTutorOffers(offerIDs, tutorID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorService) AddTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error) {
	res, err := t.client.AddTutorSubjects(tutorID, subjectIDs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *TutorService) UpdateTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error) {
	res, err := t.client.UpdateTutorSubjects(tutorID, subjectIDs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *TutorService) mapTutorsWithInfoDTO(
	tutors []tutorDto.TutorShortInfo,
	subjectsByID map[int64]subjectDto.Subject,
) ([]tutorDto.TutorShortInfoWithSubjects, error) {
	var tutorIDs []int64
	for _, oneTutor := range tutors {
		tutorIDs = append(tutorIDs, oneTutor.TutorID)
	}

	tutorsInfo, err := t.user.GetUsersShortInfo(tutorIDs)
	if err != nil {
		return nil, err
	}

	tutorsByID := make(map[int64]userDto.UserShortInfo)
	for _, oneTutor := range tutorsInfo {
		tutorsByID[oneTutor.ID] = oneTutor
	}

	var resTutors []tutorDto.TutorShortInfoWithSubjects
	for _, oneTutor := range tutors {
		var tutorSubjects []subjectDto.Subject
		for _, subjectID := range oneTutor.SubjectIDs {
			tutorSubjects = append(tutorSubjects, subjectsByID[subjectID])
		}

		resTutors = append(resTutors, tutorDto.TutorShortInfoWithSubjects{
			Tutor:         tutorsByID[oneTutor.TutorID],
			Rating:        oneTutor.Rating,
			ReviewsCount:  oneTutor.ReviewsCount,
			StudentsCount: oneTutor.StudentsCount,
			Subjects:      tutorSubjects,
		})
	}

	return resTutors, nil
}
