package services

import (
	"gitlab.com/chertokdmitry/surfapi/src/domain/images"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
)

var (
	ImageService imageServiceInterface = &imageService{}
)

type imageService struct {
}

type imageServiceInterface interface {
	Insert(i *images.Image) *errors.RestErr
}

func (ii *imageService) Insert(i *images.Image) *errors.RestErr {
	return i.Insert()
}
