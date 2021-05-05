package db_resource

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gitlab.com/chertokdmitry/surfapi/src/env"
	"gitlab.com/chertokdmitry/surfapi/src/logger"
	"gitlab.com/chertokdmitry/surfapi/src/message"
)

func GetDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		env.PG_HOST, 5432, env.PG_USER, env.PG_PASS, env.PG_DBNAME)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		logger.Error(message.ErrOpenConn, err)
	}

	err = db.Ping()

	if err != nil {
		logger.Error(message.ErrPingConn, err)
	}

	return db
}
