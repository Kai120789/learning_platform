package grpc

import (
	"context"
	tutorGRPC "github.com/Kai120789/learning_platform_proto/protos/gen/go/tutor"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-platform/tutors/internal/dto"
	"learning-platform/tutors/internal/models"
	"learning-platform/tutors/internal/utils"
)

type TutorReviewService interface {
	GetTutorReviews(getTutorReviews dto.GetTutorReviews) ([]models.TutorReview, int64, error)
	AddTutorReview(review dto.NewTutorReview) (*models.TutorReview, error)
	UpdateTutorReview(updReview dto.UpdateTutorReview) (*models.TutorReview, error)
	DeleteTutorReview(reviewID, authorID int64) error
}

func (t *TutorGRPCServer) GetTutorReviews(
	ctx context.Context,
	in *tutorGRPC.GetTutorReviewsRequest,
) (*tutorGRPC.GetTutorReviewsResponse, error) {
	request := dto.GetTutorReviews{
		TutorID: in.GetTutorId(),
		Page:    in.GetPage(),
		Limit:   in.GetLimit(),
	}

	reviews, count, err := t.TutorReviewService.GetTutorReviews(request)
	if err != nil {
		t.logger.Error(
			"failed to get tutor reviews",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get tutor reviews")

	}

	var resReviews []*tutorGRPC.Review
	for _, oneReview := range reviews {
		resReviews = append(resReviews, &tutorGRPC.Review{
			Id:        oneReview.ID,
			TutorId:   oneReview.TutorID,
			AuthorId:  oneReview.AuthorID,
			SubjectId: oneReview.SubjectID,
			Text:      oneReview.Text,
			Rating:    oneReview.Rating,
			CreatedAt: utils.DBTimeToOptional(oneReview.CreatedAt),
			UpdatedAt: utils.DBTimeToOptional(oneReview.UpdatedAt),
		})
	}

	return &tutorGRPC.GetTutorReviewsResponse{
		Reviews: resReviews,
		Count:   count,
	}, nil
}

func (t *TutorGRPCServer) AddTutorReview(
	ctx context.Context,
	in *tutorGRPC.AddTutorReviewRequest,
) (*tutorGRPC.AddTutorReviewResponse, error) {
	request := dto.NewTutorReview{
		Text:      in.GetText(),
		TutorID:   in.GetTutorId(),
		AuthorID:  in.GetAuthorId(),
		SubjectID: in.GetSubjectId(),
		Rating:    in.GetRating(),
	}

	review, err := t.TutorReviewService.AddTutorReview(request)
	if err != nil {
		t.logger.Error(
			"failed to add tutor review",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Int64("authorID", in.GetAuthorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to add tutor review")
	}

	return &tutorGRPC.AddTutorReviewResponse{
		Review: &tutorGRPC.Review{
			Id:        review.ID,
			TutorId:   review.TutorID,
			AuthorId:  review.AuthorID,
			SubjectId: review.SubjectID,
			Text:      review.Text,
			Rating:    review.Rating,
			CreatedAt: utils.DBTimeToOptional(review.CreatedAt),
			UpdatedAt: utils.DBTimeToOptional(review.UpdatedAt),
		},
	}, nil
}

func (t *TutorGRPCServer) UpdateTutorReview(
	ctx context.Context,
	in *tutorGRPC.UpdateTutorReviewRequest,
) (*tutorGRPC.UpdateTutorReviewResponse, error) {
	request := dto.UpdateTutorReview{
		ID:        in.GetId(),
		Text:      in.GetText(),
		AuthorID:  in.GetAuthorId(),
		SubjectID: in.GetSubjectId(),
		Rating:    in.GetRating(),
	}

	review, err := t.TutorReviewService.UpdateTutorReview(request)
	if err != nil {
		t.logger.Error(
			"failed to update tutor review",
			zap.Int64("reviewID", in.GetId()),
			zap.Int64("authorID", in.GetAuthorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update tutor review")
	}

	return &tutorGRPC.UpdateTutorReviewResponse{
		Review: &tutorGRPC.Review{
			Id:        review.ID,
			TutorId:   review.TutorID,
			AuthorId:  review.AuthorID,
			SubjectId: review.SubjectID,
			Text:      review.Text,
			Rating:    review.Rating,
			CreatedAt: utils.DBTimeToOptional(review.CreatedAt),
			UpdatedAt: utils.DBTimeToOptional(review.UpdatedAt),
		},
	}, nil
}

func (t *TutorGRPCServer) DeleteTutorReview(
	ctx context.Context,
	in *tutorGRPC.DeleteTutorReviewRequest,
) (*tutorGRPC.DeleteTutorReviewResponse, error) {
	err := t.TutorReviewService.DeleteTutorReview(in.GetId(), in.GetAuthorId())
	if err != nil {
		t.logger.Error(
			"failed to delete tutor review",
			zap.Int64("reviewID", in.GetId()),
			zap.Int64("authorID", in.GetAuthorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete tutor review")
	}

	return &tutorGRPC.DeleteTutorReviewResponse{}, nil
}
