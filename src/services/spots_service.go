package services

import (
	"gitlab.com/chertokdmitry/surfapi/src/domain/spots"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
)

var (
	SpotsService spotsServiceInterface = &spotsService{}
)

type spotsService struct {
}

type spotsServiceInterface interface {
	GetByRegionId(int64) ([]spots.Spot, *errors.RestErr)
	GetAll() ([]spots.Spot, *errors.RestErr)
}

func (ss *spotsService) GetByRegionId(regionId int64) ([]spots.Spot, *errors.RestErr) {
	s := &spots.Spot{}
	result, err := s.GetByRegionId(regionId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ss *spotsService) GetAll() ([]spots.Spot, *errors.RestErr) {
	s := &spots.Spot{}
	result, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}
