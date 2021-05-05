package weather

import (
	"gitlab.com/chertokdmitry/surfapi/src/db_resource/weather_db"
	"gitlab.com/chertokdmitry/surfapi/src/logger"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
)

const querySelectWeatherBySpot = "SELECT weather.id, spots.title, weather.hourly FROM spots LEFT JOIN  weather ON spots.id =  weather.spot_id WHERE spots.id = $1 ORDER BY weather.created_at DESC limit 1"

func (weather *Weather) Get() *errors.RestErr {
	db := db_resource.GetDB()
	defer db.Close()

	result := db.QueryRow(querySelectWeatherBySpot, weather.SpotId)
	if err := result.Scan(&weather.SpotId, &weather.Title, &weather.Hourly); err != nil {
		logger.Error("error when trying to get weather", err)
		return errors.NewInternalServerError("error when trying to get weather")
	}

	return nil
}
