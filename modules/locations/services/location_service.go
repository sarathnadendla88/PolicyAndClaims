package location_service

import (
	location "policy/modules/locations/domain"
	"policy/utils/date_utils"
	"policy/utils/errors"
	"policy/utils/utils"
)

var (
	LocationService userServiceInterface = &locationService{}
)

type locationService struct {
}

type userServiceInterface interface {
	CreateLocation(location location.Location) (*location.Location, *errors.RestErr)
	GetAllLocations() ([]location.Location, *errors.RestErr)
	GetLocationById(string) (*location.Location, *errors.RestErr)
}

func (s *locationService) GetAllLocations() ([]location.Location, *errors.RestErr) {
	result := &location.Location{}
	return result.GetAllLocations()
}

func (s *locationService) GetLocationById(id string) (*location.Location, *errors.RestErr) {
	result := &location.Location{Id: id}
	err := result.GetLocationById()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *locationService) CreateLocation(location location.Location) (*location.Location, *errors.RestErr) {
	id, err := utils.GetRandomString()
	if err != nil {
		return nil, err
	}
	location.Id = id
	date := date_utils.GetNowStringForDB()

	location.CreatedDateTime = &date
	if err := location.Save(); err != nil {
		return nil, err
	}
	return &location, nil
}
