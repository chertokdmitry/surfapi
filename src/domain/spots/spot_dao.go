package spots

import (
	"gitlab.com/chertokdmitry/surfapi/src/db_resource/weather_db"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
)

const querySelectSpotsByRegion = "SELECT id, title FROM spots WHERE region_id = $1"
const querySelectSpotsAll = "SELECT id, title FROM spots ORDER BY id"

var SpotIds []int

func (s *Spot) GetByRegionId(regionId int64) ([]Spot, *errors.RestErr) {
	db := db_resource.GetDB()
	defer db.Close()

	rows, err := db.Query(querySelectSpotsByRegion, regionId)
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get spots")
	}
	defer rows.Close()

	spot := Spot{}
	var res []Spot
	for rows.Next() {
		err = rows.Scan(&spot.Id, &spot.Title)
		if err != nil {
			return nil, errors.NewInternalServerError("error when scan spots")
		}

		res = append(res, spot)
	}

	return res, nil
}

func (s *Spot) GetAll() ([]Spot, *errors.RestErr) {
	db := db_resource.GetDB()
	defer db.Close()

	rows, err := db.Query(querySelectSpotsAll)
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get all spots")
	}
	defer rows.Close()

	spot := Spot{}
	var res []Spot
	for rows.Next() {
		err = rows.Scan(&spot.Id, &spot.Title)
		if err != nil {
			return nil, errors.NewInternalServerError("error when scan all spots")
		}

		res = append(res, spot)
	}

	return res, nil
}
