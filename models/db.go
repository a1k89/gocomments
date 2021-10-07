package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		panic("Provide .env file")
	}
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPassword)

	connection, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic("Can't connect to db")
	}

	db = connection
	db.Debug().AutoMigrate(&Comment{}, &User{})
}

func GetDB() *gorm.DB {
	return db
}
