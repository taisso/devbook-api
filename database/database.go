package database

import (
	"api/database/migrations"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB(str string) {
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatal("error: ", err)
	}

	db = database

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
