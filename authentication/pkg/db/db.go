package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func ConnectDB(name string) *gorm.DB {

	db, err := gorm.Open(
		sqlite.Open(name),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	
	if err != nil {
		log.Fatal(err)
	}

	return db
}
