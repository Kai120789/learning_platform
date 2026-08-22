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

type TutorOfferService interface {
	GetTutorOffers(tutorID int64) ([]models.TutorOffer, error)
	AddTutorOffer(offer dto.NewTutorOffer) (*models.TutorOffer, error)
	UpdateTutorOffer(updOffer dto.UpdateTutorOffer) (*models.TutorOffer, error)
	DeleteOneTutorOffer(offerID, tutorID int64) error
	DeleteTutorOffers(offerIDs []int64, tutorID int64) error
}

func (t *TutorGRPCServer) GetTutorOffers(
	ctx context.Context,
	in *tutorGRPC.GetTutorOffersRequest,
) (*tutorGRPC.GetTutorOffersResponse, error) {
	offers, err := t.TutorOfferService.GetTutorOffers(in.GetTutorId())
	if err != nil {
		t.logger.Error(
			"failed to get tutor offers",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to get tutor offers")
	}

	var resOffers []*tutorGRPC.Offer
	for _, oneOffer := range offers {
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

	return &tutorGRPC.GetTutorOffersResponse{
		Offers: resOffers,
	}, nil
}

func (t *TutorGRPCServer) AddTutorOffer(
	ctx context.Context,
	in *tutorGRPC.AddTutorOfferRequest,
) (*tutorGRPC.AddTutorOfferResponse, error) {
	request := dto.NewTutorOffer{
		TutorID:         in.GetTutorId(),
		SubjectID:       in.GetSubjectId(),
		Title:           in.GetTitle(),
		Description:     in.Description,
		Price:           in.GetPrice(),
		DurationMinutes: in.DurationMinutes,
	}

	offer, err := t.TutorOfferService.AddTutorOffer(request)
	if err != nil {
		t.logger.Error(
			"failed to add tutor offer",
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to add tutor offer")
	}

	return &tutorGRPC.AddTutorOfferResponse{
		Offer: &tutorGRPC.Offer{
			Id:              offer.ID,
			Title:           offer.Title,
			Description:     utils.DBStringToOptional(offer.Description),
			SubjectId:       offer.SubjectID,
			TutorId:         offer.TutorID,
			Price:           offer.Price,
			DurationMinutes: utils.DBInt8ToOptional(offer.DurationMinutes),
		},
	}, nil
}

func (t *TutorGRPCServer) UpdateTutorOffer(
	ctx context.Context,
	in *tutorGRPC.UpdateTutorOfferRequest,
) (*tutorGRPC.UpdateTutorOfferResponse, error) {
	request := dto.UpdateTutorOffer{
		ID:              in.GetId(),
		TutorID:         in.GetTutorId(),
		SubjectID:       in.GetSubjectId(),
		Title:           in.GetTitle(),
		Description:     in.Description,
		Price:           in.GetPrice(),
		DurationMinutes: in.DurationMinutes,
	}

	offer, err := t.TutorOfferService.UpdateTutorOffer(request)
	if err != nil {
		t.logger.Error(
			"failed to update tutor offer",
			zap.Int64("offerID", in.GetId()),
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to update tutor offer")
	}

	return &tutorGRPC.UpdateTutorOfferResponse{
		Offer: &tutorGRPC.Offer{
			Id:              offer.ID,
			Title:           offer.Title,
			Description:     utils.DBStringToOptional(offer.Description),
			SubjectId:       offer.SubjectID,
			TutorId:         offer.TutorID,
			Price:           offer.Price,
			DurationMinutes: utils.DBInt8ToOptional(offer.DurationMinutes),
		},
	}, nil
}

func (t *TutorGRPCServer) DeleteOneTutorOffer(
	ctx context.Context,
	in *tutorGRPC.DeleteOneTutorOfferRequest,
) (*tutorGRPC.DeleteOneTutorOfferResponse, error) {
	err := t.TutorOfferService.DeleteOneTutorOffer(in.GetId(), in.GetTutorId())
	if err != nil {
		t.logger.Error(
			"failed to delete one tutor offer",
			zap.Int64("offerID", in.GetId()),
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete one tutor offer")
	}

	return &tutorGRPC.DeleteOneTutorOfferResponse{}, nil
}

func (t *TutorGRPCServer) DeleteTutorOffers(
	ctx context.Context,
	in *tutorGRPC.DeleteTutorOffersRequest,
) (*tutorGRPC.DeleteTutorOffersResponse, error) {
	err := t.TutorOfferService.DeleteTutorOffers(in.GetIds(), in.GetTutorId())
	if err != nil {
		t.logger.Error(
			"failed to delete tutor offers",
			zap.Int64s("offerIDs", in.GetIds()),
			zap.Int64("tutorID", in.GetTutorId()),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, "failed to delete tutor offers")
	}

	return &tutorGRPC.DeleteTutorOffersResponse{}, nil
}
