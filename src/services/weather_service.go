package services

import (
	"gitlab.com/chertokdmitry/surfapi/src/domain/weather"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
)

var (
	WeatherService weatherServiceInterface = &weatherService{}
)

type weatherService struct {
}

type weatherServiceInterface interface {
	GetWeather(int64) (*weather.Weather, *errors.RestErr)
}

func (w *weatherService) GetWeather(spotId int64) (*weather.Weather, *errors.RestErr) {
	result := &weather.Weather{SpotId: spotId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}
