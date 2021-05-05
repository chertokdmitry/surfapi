package users

import (
	"fmt"
	_ "github.com/lib/pq"
	db_resource "gitlab.com/chertokdmitry/surfapi/src/db_resource/weather_db"
	"gitlab.com/chertokdmitry/surfapi/src/logger"
	"gitlab.com/chertokdmitry/surfapi/src/utils/date_utils"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
)

const (
	queryInsertUser       = `INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES ($1, $2, $3, $4, $5, $6)`
	queryGetUser          = `SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = $1`
	queryUpdateUser       = `UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4`
	queryDeleteUser       = `DELETE FROM users WHERE id= $1`
	queryFindUserByStatus = `SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = $1`
)

func (user *User) Get() *errors.RestErr {
	db := db_resource.GetDB()

	defer db.Close()

	result := db.QueryRow(queryGetUser, user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to get user", err)
		return errors.NewInternalServerError("error when trying to get user")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	db := db_resource.GetDB()
	defer db.Close()

	_, err := db.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, date_utils.GetNow(), user.Status, user.Password)

	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("error when trying to save user")
	}

	return nil
}

func (user *User) Update() *errors.RestErr {
	db := db_resource.GetDB()
	defer db.Close()
	_, err := db.Exec(queryUpdateUser, user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("update error", err)
		return errors.NewInternalServerError("update error")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	db := db_resource.GetDB()
	defer db.Close()

	_, err := db.Exec(queryDeleteUser, user.Id)
	if err != nil {
		logger.Error("delete error", err)
		return errors.NewInternalServerError("delete error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	db := db_resource.GetDB()

	rows, err := db.Query(queryFindUserByStatus, status)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("FindByStatus error", err)
			return nil, errors.NewInternalServerError("FindByStatus error")
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
