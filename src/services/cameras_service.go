package services

import (
	"gitlab.com/chertokdmitry/surfapi/src/domain/cameras"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
)

var (
	CamerasService camerasServiceInterface = &camerasService{}
)

type camerasService struct {
}

type camerasServiceInterface interface {
	GetAll() ([]int64, *errors.RestErr)
	GetBySpot(int64) ([]int64, *errors.RestErr)
}

func (cs *camerasService) GetAll() ([]int64, *errors.RestErr) {
	c := &cameras.Camera{}
	result, err := c.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *camerasService) GetBySpot(id int64) ([]int64, *errors.RestErr) {
	c := &cameras.Camera{}
	result, err := c.GetBySpot(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
