package grpc

import (
	"context"
	tutorGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/tutor"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/tutors/internal/dto"
	"learning-platform/tutors/internal/utils"
)

type TutorBaseService interface {
	GetOneTutor(tutorID, userID int64) (*dto.OneTutor, error)
	GetTutors(getTutors dto.GetTutors) ([]dto.TutorShortInfo, int64, error)
}

func (t *TutorGRPCServer) GetOneTutor(
	ctx context.Context,
	in *tutorGRPC.GetOneTutorRequest,
) (*tutorGRPC.GetOneTutorResponse, error) {
	tutor, err := t.TutorBaseService.GetOneTutor(in.GetTutorId(), in.GetUserId())
	if err != nil {
		t.logger.Error(
			"failed to get one tutor",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Int64("userID", in.GetUserId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get one tutor")
	}

	var resOffers []*tutorGRPC.Offer
	for _, oneOffer := range tutor.Offers {
		resOffers = append(resOffers, &tutorGRPC.Offer{
			Id:              oneOffer.ID,
			Title:           oneOffer.Title,
			Description:     utils.DBStringToOptional(oneOffer.Description),
			SubjectId:       oneOffer.SubjectID,
			TutorId:         oneOffer.TutorID,
			Price:           oneOffer.Price,
			DurationMinutes: utils.DBInt8ToOptional(oneOffer.DurationMinutes),
		})
	}

	var myReview *tutorGRPC.Review
	if tutor.MyReview != nil {
		myReview = &tutorGRPC.Review{
			Id:        tutor.MyReview.ID,
			TutorId:   tutor.MyReview.TutorID,
			AuthorId:  tutor.MyReview.AuthorID,
			SubjectId: tutor.MyReview.SubjectID,
			Text:      tutor.MyReview.Text,
			Rating:    tutor.MyReview.Rating,
			CreatedAt: utils.DBTimeToOptional(tutor.MyReview.CreatedAt),
			UpdatedAt: utils.DBTimeToOptional(tutor.MyReview.UpdatedAt),
		}
	}

	return &tutorGRPC.GetOneTutorResponse{
		TutorInfo: &tutorGRPC.TutorShortInfo{
			TutorId:       tutor.TutorInfo.TutorID,
			Rating:        tutor.TutorInfo.Rating,
			ReviewsCount:  tutor.TutorInfo.ReviewsCount,
			StudentsCount: tutor.TutorInfo.StudentsCount,
			SubjectIds:    tutor.TutorInfo.SubjectIDs,
		},
		Offers:   resOffers,
		MyReview: myReview,
	}, nil
}

func (t *TutorGRPCServer) GetTutors(
	ctx context.Context,
	in *tutorGRPC.GetTutorsRequest,
) (*tutorGRPC.GetTutorsResponse, error) {
	tutors, count, err := t.TutorBaseService.GetTutors(dto.GetTutors{
		SubjectID: in.SubjectId,
		Page:      in.GetPage(),
		Limit:     in.GetLimit(),
	})
	if err != nil {
		t.logger.Error(
			"failed to get tutors",
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get tutors")
	}

	var resTutors []*tutorGRPC.TutorShortInfo
	for _, oneTutor := range tutors {
		resTutors = append(resTutors, &tutorGRPC.TutorShortInfo{
			TutorId:       oneTutor.TutorID,
			Rating:        oneTutor.Rating,
			ReviewsCount:  oneTutor.ReviewsCount,
			StudentsCount: oneTutor.StudentsCount,
			SubjectIds:    oneTutor.SubjectIDs,
		})
	}

	return &tutorGRPC.GetTutorsResponse{
		Tutors: resTutors,
		Count:  count,
	}, nil
}
