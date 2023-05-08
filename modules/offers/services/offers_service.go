package offer_service

import (
	offer "policy/modules/offers/domain"
	"policy/utils/date_utils"
	"policy/utils/errors"
	"policy/utils/utils"
)

var (
	OfferService offerServiceInterface = &offerService{}
)

type offerService struct {
}

type offerServiceInterface interface {
	CreateOffer(offer offer.Offers) (*offer.Offers, *errors.RestErr)
	GetAllOffers() ([]offer.Offers, *errors.RestErr)
	GetOfferById(string) (*offer.Offers, *errors.RestErr)
}

func (s *offerService) GetAllOffers() ([]offer.Offers, *errors.RestErr) {
	result := &offer.Offers{}
	return result.GetAllOffers()
}

func (s *offerService) GetOfferById(id string) (*offer.Offers, *errors.RestErr) {
	result := &offer.Offers{Id: id}
	err := result.GetOfferById()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *offerService) CreateOffer(offers offer.Offers) (*offer.Offers, *errors.RestErr) {
	id, err := utils.GetRandomString()
	if err != nil {
		return nil, err
	}
	offers.Id = id
	date := date_utils.GetNowStringForDB()

	offers.CreatedDateTime = &date
	if err := offers.Save(); err != nil {
		return nil, err
	}
	return &offers, nil
}
