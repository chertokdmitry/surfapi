package subscriptions

import (
	"gitlab.com/chertokdmitry/surfapi/src/db_resource/weather_db"
	"gitlab.com/chertokdmitry/surfapi/src/logger"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
	"time"
)

const (
	queryInsertNewSub       = "INSERT INTO subscription(chat_id, spot_id, hour, created_at) VALUES ($1, $2, $3, $4)"
	querySelectSubsByChatId = "SELECT subscription.id, spots.title, subscription.hour FROM subscription LEFT JOIN spots ON subscription.spot_id = spots.id WHERE chat_id = $1 ORDER BY subscription.created_at"
	queryDeleteSub          = "DELETE FROM subscription WHERE id = $1"
)

// insert new sub
func (s *Sub) Insert() *errors.RestErr {
	db := db_resource.GetDB()
	defer db.Close()

	_, err := db.Exec(queryInsertNewSub, s.ChatId, s.SpotId, s.Hour, time.Now())
	if err != nil {
		logger.Error("insert error", err)
		return errors.NewInternalServerError("insert error")
	}

	return nil
}

// delete sub
func (s *Sub) Delete() *errors.RestErr {
	db := db_resource.GetDB()
	defer db.Close()

	_, err := db.Exec(queryDeleteSub, s.Id)
	if err != nil {
		logger.Error("delete error", err)
		return errors.NewInternalServerError("delete error")
	}
	return nil
}

// get subs by chat id
func (s *Sub) GetByChatId(chatId int64) ([]Sub, *errors.RestErr) {
	db := db_resource.GetDB()
	defer db.Close()

	rows, err := db.Query(querySelectSubsByChatId, chatId)
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get subs by chat id")
	}
	defer rows.Close()

	sub := Sub{}
	var res []Sub
	for rows.Next() {
		err = rows.Scan(&sub.Id, &sub.Title, &sub.Hour)
		if err != nil {
			return nil, errors.NewInternalServerError("error when scan subs")
		}

		res = append(res, sub)
	}

	return res, nil
}
