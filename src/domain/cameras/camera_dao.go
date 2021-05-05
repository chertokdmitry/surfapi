package cameras

import (
	"gitlab.com/chertokdmitry/surfapi/src/db_resource/weather_db"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
)

const querySelectCamerasAll = "SELECT id FROM cameras"
const querySelectCamerasBySpot = "SELECT camera_id FROM spot_cameras WHERE spot_id = $1"

var CameraIds []int

func (c *Camera) GetAll() ([]int64, *errors.RestErr) {
	db := db_resource.GetDB()
	defer db.Close()

	rows, err := db.Query(querySelectCamerasAll)
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get all cameras")
	}
	defer rows.Close()

	camera := Camera{}
	var res []int64
	for rows.Next() {
		err = rows.Scan(&camera.Id)
		if err != nil {
			return nil, errors.NewInternalServerError("error when scan all camera")
		}

		res = append(res, camera.Id)
	}

	return res, nil
}

func (c *Camera) GetBySpot(id int64) ([]int64, *errors.RestErr) {
	db := db_resource.GetDB()
	defer db.Close()

	rows, err := db.Query(querySelectCamerasBySpot, id)
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get cameras")
	}
	defer rows.Close()

	var cameraId int64
	var res []int64

	for rows.Next() {
		err = rows.Scan(&cameraId)

		if err != nil {
			return nil, errors.NewInternalServerError("error when scan cameras")
		}

		res = append(res, cameraId)
	}

	return res, nil
}
