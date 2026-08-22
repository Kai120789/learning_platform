package service

import (
	"learning-platform/tutors/internal/dto"
	"learning-platform/tutors/internal/models"
)

type TutorOfferService struct {
	storage TutorOfferStorage
}

type TutorOfferStorage interface {
	GetTutorOffers(tutorID int64) ([]models.TutorOffer, error)
	AddTutorOffer(offer dto.NewTutorOffer) (*models.TutorOffer, error)
	UpdateTutorOffer(updOffer dto.UpdateTutorOffer) (*models.TutorOffer, error)
	DeleteOneTutorOffer(offerID, tutorID int64) error
	DeleteTutorOffers(offerIDs []int64, tutorID int64) error
}

func NewTutorOfferService(storage TutorOfferStorage) *TutorOfferService {
	return &TutorOfferService{
		storage: storage,
	}
}

func (to *TutorOfferService) GetTutorOffers(tutorID int64) ([]models.TutorOffer, error) {
	res, err := to.storage.GetTutorOffers(tutorID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (to *TutorOfferService) AddTutorOffer(offer dto.NewTutorOffer) (*models.TutorOffer, error) {
	res, err := to.storage.AddTutorOffer(offer)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (to *TutorOfferService) UpdateTutorOffer(updOffer dto.UpdateTutorOffer) (*models.TutorOffer, error) {
	res, err := to.storage.UpdateTutorOffer(updOffer)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (to *TutorOfferService) DeleteOneTutorOffer(offerID, tutorID int64) error {
	err := to.storage.DeleteOneTutorOffer(offerID, tutorID)
	if err != nil {
		return err
	}

	return nil
}

func (to *TutorOfferService) DeleteTutorOffers(offerIDs []int64, tutorID int64) error {
	err := to.storage.DeleteTutorOffers(offerIDs, tutorID)
	if err != nil {
		return err
	}

	return nil
}
