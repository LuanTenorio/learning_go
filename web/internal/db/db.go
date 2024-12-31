package db

import (
	"database/sql"
)

func ConnectDB(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err.Error())
	}

	return db
}
