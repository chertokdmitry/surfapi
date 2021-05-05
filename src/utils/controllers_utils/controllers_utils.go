package controllers_utils

import (
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
	"strconv"
)

func CheckId(idParam string) (int64, *errors.RestErr) {
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("ID should be number")
	}

	return id, nil
}
