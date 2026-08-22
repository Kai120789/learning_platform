package grpc

import (
	"context"
	tutorGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/tutor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learning-platform/api-gateway/internal/dto/tutorDto"
	"time"
)

type TutorClient struct {
	client tutorGRPC.TutorClient
}

func NewTutorGrpcConnection(tutorGrpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		tutorGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewTutorClient(connection *grpc.ClientConn) *TutorClient {
	return &TutorClient{
		client: tutorGRPC.NewTutorClient(connection),
	}
}

func (t *TutorClient) AddStudent(tutorID, studentID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.client.AddStudent(ctx, &tutorGRPC.AddStudentRequest{
		TutorId:   tutorID,
		StudentId: studentID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorClient) DeleteOneStudent(tutorID, studentID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.client.DeleteOneStudent(ctx, &tutorGRPC.DeleteOneStudentRequest{
		TutorId:   tutorID,
		StudentId: studentID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorClient) DeleteStudents(tutorID int64, studentIDs []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.client.DeleteStudents(ctx, &tutorGRPC.DeleteStudentsRequest{
		TutorId:    tutorID,
		StudentIds: studentIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorClient) GetTutorStudents(getTutorStudents tutorDto.GetTutorStudents) ([]tutorDto.TutorStudent, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	studentsWithCount, err := t.client.GetTutorStudents(ctx, &tutorGRPC.GetTutorStudentsRequest{
		TutorId:              getTutorStudents.TutorID,
		InteractedWithinDays: getTutorStudents.InteractedWithinDays,
		Page:                 getTutorStudents.Page,
		Limit:                getTutorStudents.Limit,
	})
	if err != nil {
		return nil, 0, err
	}

	var resTutorStudents []tutorDto.TutorStudent
	for _, oneTutorStudent := range studentsWithCount.GetStudents() {
		var lastInteractedAt *time.Time
		if oneTutorStudent.LastInteractedAt != nil {
			lia := oneTutorStudent.LastInteractedAt.AsTime()
			lastInteractedAt = &lia
		}

		resTutorStudents = append(resTutorStudents, tutorDto.TutorStudent{
			StudentID:        oneTutorStudent.GetStudentId(),
			LastInteractedAt: lastInteractedAt,
		})
	}

	return resTutorStudents, studentsWithCount.Count, nil
}

func (t *TutorClient) GetOneTutor(tutorID, userID int64) (*tutorDto.OneTutor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tutor, err := t.client.GetOneTutor(ctx, &tutorGRPC.GetOneTutorRequest{
		TutorId: tutorID,
		UserId:  userID,
	})
	if err != nil {
		return nil, err
	}

	var resOffers []tutorDto.TutorOffer
	for _, oneOffer := range tutor.GetOffers() {
		resOffers = append(resOffers, tutorDto.TutorOffer{
			ID:              oneOffer.GetId(),
			TutorID:         oneOffer.GetTutorId(),
			SubjectID:       oneOffer.GetSubjectId(),
			Title:           oneOffer.GetTitle(),
			Price:           oneOffer.Price,
			Description:     oneOffer.Description,
			DurationMinutes: oneOffer.DurationMinutes,
		})
	}

	var myReview *tutorDto.TutorReview
	if tutor.MyReview != nil {
		myReview = &tutorDto.TutorReview{
			ID:        tutor.GetMyReview().GetId(),
			TutorID:   tutor.GetMyReview().GetTutorId(),
			AuthorID:  tutor.GetMyReview().GetAuthorId(),
			SubjectID: tutor.GetMyReview().GetSubjectId(),
			Text:      tutor.GetMyReview().GetText(),
			Rating:    tutor.GetMyReview().GetRating(),
			CreatedAt: tutor.GetMyReview().GetCreatedAt().AsTime(),
			UpdatedAt: tutor.GetMyReview().GetUpdatedAt().AsTime(),
		}
	}

	return &tutorDto.OneTutor{
		TutorInfo: tutorDto.TutorShortInfo{
			TutorID:       tutor.GetTutorInfo().GetTutorId(),
			Rating:        tutor.GetTutorInfo().GetRating(),
			ReviewsCount:  tutor.GetTutorInfo().GetReviewsCount(),
			StudentsCount: tutor.GetTutorInfo().GetStudentsCount(),
			SubjectIDs:    tutor.GetTutorInfo().GetSubjectIds(),
		},
		Offers:   resOffers,
		MyReview: myReview,
	}, nil
}

func (t *TutorClient) GetTutors(getTutors tutorDto.GetTutors) ([]tutorDto.TutorShortInfo, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tutorsWithCount, err := t.client.GetTutors(ctx, &tutorGRPC.GetTutorsRequest{
		SubjectId: getTutors.SubjectID,
		Page:      getTutors.Page,
		Limit:     getTutors.Limit,
	})
	if err != nil {
		return nil, 0, err
	}

	var resTutors []tutorDto.TutorShortInfo
	for _, oneTutor := range tutorsWithCount.GetTutors() {
		resTutors = append(resTutors, tutorDto.TutorShortInfo{
			TutorID:       oneTutor.GetTutorId(),
			Rating:        oneTutor.GetRating(),
			ReviewsCount:  oneTutor.GetReviewsCount(),
			StudentsCount: oneTutor.GetStudentsCount(),
			SubjectIDs:    oneTutor.GetSubjectIds(),
		})
	}

	return resTutors, tutorsWithCount.GetCount(), nil
}

func (t *TutorClient) GetTutorReviews(getTutorReviews tutorDto.GetTutorReviews) ([]tutorDto.TutorReview, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reviewsWithCount, err := t.client.GetTutorReviews(ctx, &tutorGRPC.GetTutorReviewsRequest{
		TutorId: getTutorReviews.TutorID,
		Page:    getTutorReviews.Page,
		Limit:   getTutorReviews.Limit,
	})
	if err != nil {
		return nil, 0, err
	}

	var resReviews []tutorDto.TutorReview
	for _, oneReview := range reviewsWithCount.GetReviews() {
		resReviews = append(resReviews, tutorDto.TutorReview{
			ID:        oneReview.GetId(),
			TutorID:   oneReview.GetTutorId(),
			AuthorID:  oneReview.GetAuthorId(),
			SubjectID: oneReview.GetSubjectId(),
			Text:      oneReview.GetText(),
			Rating:    oneReview.GetRating(),
			CreatedAt: oneReview.GetCreatedAt().AsTime(),
			UpdatedAt: oneReview.GetUpdatedAt().AsTime(),
		})
	}

	return resReviews, reviewsWithCount.GetCount(), nil
}

func (t *TutorClient) AddTutorReview(newReview tutorDto.NewTutorReview) (*tutorDto.TutorReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	review, err := t.client.AddTutorReview(ctx, &tutorGRPC.AddTutorReviewRequest{
		TutorId:   newReview.TutorID,
		AuthorId:  newReview.AuthorID,
		SubjectId: newReview.SubjectID,
		Text:      newReview.Text,
		Rating:    newReview.Rating,
	})
	if err != nil {
		return nil, err
	}

	return &tutorDto.TutorReview{
		ID:        review.GetReview().GetId(),
		TutorID:   review.GetReview().GetTutorId(),
		AuthorID:  review.GetReview().GetAuthorId(),
		SubjectID: review.GetReview().GetSubjectId(),
		Text:      review.GetReview().GetText(),
		Rating:    review.GetReview().GetRating(),
		CreatedAt: review.GetReview().GetCreatedAt().AsTime(),
		UpdatedAt: review.GetReview().GetUpdatedAt().AsTime(),
	}, nil
}

func (t *TutorClient) UpdateTutorReview(updReview tutorDto.UpdateTutorReview) (*tutorDto.TutorReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	review, err := t.client.UpdateTutorReview(ctx, &tutorGRPC.UpdateTutorReviewRequest{
		Id:        updReview.ID,
		AuthorId:  updReview.AuthorID,
		SubjectId: updReview.SubjectID,
		Text:      updReview.Text,
		Rating:    updReview.Rating,
	})
	if err != nil {
		return nil, err
	}

	return &tutorDto.TutorReview{
		ID:        review.GetReview().GetId(),
		TutorID:   review.GetReview().GetTutorId(),
		AuthorID:  review.GetReview().GetAuthorId(),
		SubjectID: review.GetReview().GetSubjectId(),
		Text:      review.GetReview().GetText(),
		Rating:    review.GetReview().GetRating(),
		CreatedAt: review.GetReview().GetCreatedAt().AsTime(),
		UpdatedAt: review.GetReview().GetUpdatedAt().AsTime(),
	}, nil
}

func (t *TutorClient) DeleteTutorReview(reviewID, authorID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.client.DeleteTutorReview(ctx, &tutorGRPC.DeleteTutorReviewRequest{
		Id:       reviewID,
		AuthorId: authorID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorClient) GetTutorOffers(tutorID int64) ([]tutorDto.TutorOffer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	offers, err := t.client.GetTutorOffers(ctx, &tutorGRPC.GetTutorOffersRequest{TutorId: tutorID})
	if err != nil {
		return nil, err
	}

	var resOffers []tutorDto.TutorOffer
	for _, oneOffer := range offers.GetOffers() {
		resOffers = append(resOffers, tutorDto.TutorOffer{
			ID:              oneOffer.GetId(),
			TutorID:         oneOffer.GetTutorId(),
			SubjectID:       oneOffer.GetSubjectId(),
			Title:           oneOffer.GetTitle(),
			Price:           oneOffer.Price,
			Description:     oneOffer.Description,
			DurationMinutes: oneOffer.DurationMinutes,
		})
	}

	return resOffers, nil
}

func (t *TutorClient) AddTutorOffer(newOffer tutorDto.NewTutorOffer) (*tutorDto.TutorOffer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	offer, err := t.client.AddTutorOffer(ctx, &tutorGRPC.AddTutorOfferRequest{
		Title:           newOffer.Title,
		Description:     newOffer.Description,
		TutorId:         newOffer.TutorID,
		SubjectId:       newOffer.SubjectID,
		Price:           newOffer.Price,
		DurationMinutes: newOffer.DurationMinutes,
	})
	if err != nil {
		return nil, err
	}

	return &tutorDto.TutorOffer{
		ID:              offer.GetOffer().GetId(),
		TutorID:         offer.GetOffer().GetTutorId(),
		SubjectID:       offer.GetOffer().GetSubjectId(),
		Title:           offer.GetOffer().GetTitle(),
		Price:           offer.GetOffer().Price,
		Description:     offer.GetOffer().Description,
		DurationMinutes: offer.GetOffer().DurationMinutes,
	}, nil
}

func (t *TutorClient) UpdateTutorOffer(updOffer tutorDto.UpdateTutorOffer) (*tutorDto.TutorOffer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	offer, err := t.client.UpdateTutorOffer(ctx, &tutorGRPC.UpdateTutorOfferRequest{
		Id:              updOffer.ID,
		Title:           updOffer.Title,
		Description:     updOffer.Description,
		TutorId:         updOffer.TutorID,
		SubjectId:       updOffer.SubjectID,
		Price:           updOffer.Price,
		DurationMinutes: updOffer.DurationMinutes,
	})
	if err != nil {
		return nil, err
	}

	return &tutorDto.TutorOffer{
		ID:              offer.GetOffer().GetId(),
		TutorID:         offer.GetOffer().GetTutorId(),
		SubjectID:       offer.GetOffer().GetSubjectId(),
		Title:           offer.GetOffer().GetTitle(),
		Description:     offer.GetOffer().Description,
		Price:           offer.GetOffer().Price,
		DurationMinutes: offer.GetOffer().DurationMinutes,
	}, nil
}

func (t *TutorClient) DeleteOneTutorOffer(offerID, tutorID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.client.DeleteOneTutorOffer(ctx, &tutorGRPC.DeleteOneTutorOfferRequest{
		Id:      offerID,
		TutorId: tutorID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorClient) DeleteTutorOffers(offerIDs []int64, tutorID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.client.DeleteTutorOffers(ctx, &tutorGRPC.DeleteTutorOffersRequest{
		Ids:     offerIDs,
		TutorId: tutorID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *TutorClient) AddTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resSubjectIDs, err := t.client.AddTutorSubjects(ctx, &tutorGRPC.AddTutorSubjectsRequest{
		TutorId:    tutorID,
		SubjectIds: subjectIDs,
	})
	if err != nil {
		return nil, err
	}

	return resSubjectIDs.GetSubjectIds(), nil
}

func (t *TutorClient) UpdateTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resSubjectIDs, err := t.client.UpdateTutorSubjects(ctx, &tutorGRPC.UpdateTutorSubjectsRequest{
		TutorId:    tutorID,
		SubjectIds: subjectIDs,
	})
	if err != nil {
		return nil, err
	}

	return resSubjectIDs.GetSubjectIds(), nil
}
