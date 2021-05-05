package images

import (
	"gitlab.com/chertokdmitry/surfapi/src/db_resource/weather_db"
	"gitlab.com/chertokdmitry/surfapi/src/logger"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
	"time"
)

const (
	queryInsertNewImage = "INSERT INTO images(camera_id, name, created_at) VALUES ($1, $2, $3)"
)

// insert new image
func (i *Image) Insert() *errors.RestErr {
	db := db_resource.GetDB()
	defer db.Close()
	_, err := db.Exec(queryInsertNewImage, i.CameraId, i.Name, time.Now())
	if err != nil {
		logger.Error("insert error", err)
		return errors.NewInternalServerError("insert error")
	}

	return nil
}
