package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func NewDatabase() (db *gorm.DB, err error) {
	err = godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading.env")
	}
	db, err = gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD"),
		),
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return
}
