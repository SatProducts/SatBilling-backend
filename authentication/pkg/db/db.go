package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func ConnectDB() *gorm.DB {

	dsn := "host=db dbname=users port=5432 sslmode=disable user=postgres password=" + os.Getenv("DB_PASSWORD")

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	
	if err != nil {
		log.Fatal(err)
	}

	return db
}
