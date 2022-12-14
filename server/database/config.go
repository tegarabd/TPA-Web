package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func GetDB() (*gorm.DB) {
	return db
}

func CloseDB() {
	postgresDB, _ := db.DB()
	postgresDB.Close()
}

func Connect() {

	var host string

	if os.Getenv("STATUS") == "PROD" {
		host = os.Getenv("DB_HOST_PROD")
	} else {
		host = os.Getenv("DB_HOST_DEV")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", 
						host,
						os.Getenv("DB_USER"),
						os.Getenv("DB_PASSWORD"),
						os.Getenv("DB_NAME"),
						os.Getenv("DB_PORT"))
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error connect to database")
	}
}