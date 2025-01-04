package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	err    error
	logger = log.New(os.Stdout, "DB: ", 2)
)

func DatabaseConnect() {
	connectionStr := "postgres://root:root@localhost:15432/root"

	DB, err = gorm.Open(postgres.Open(connectionStr))

	if err != nil {
		logger.Panic(err)
	}
}
