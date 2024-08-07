package database

import (
	"echommerce/internal/models/users_model"

	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// initialize connection to database

func Init(host, user, password, dbName, port string) *gorm.DB {
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable`, host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err := db.AutoMigrate(&users_model.User{}); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	DB = db

	return DB
}
