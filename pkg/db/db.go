package db

import (
	"log"

	// "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {

	// For postgres
	/*
		dsn := ""

		db, err := gorm.Open(
			postgres.Open(dsn),
			&gorm.Config{
				Logger: logger.Default.LogMode(logger.Silent),
			},
		)
	*/

	db, err := gorm.Open(
		sqlite.Open(
			"podbilling.db",
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
