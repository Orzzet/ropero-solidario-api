package database

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func NewDatabase() (db *gorm.DB, err error) {
	staticFiles := packr.New("static", "../../static")
	s, err := staticFiles.FindString(".env")
	if err != nil {
		log.Fatalf(err.Error())
	}
	env, err := godotenv.Unmarshal(s)
	if err != nil {
		log.Println(env)
		log.Fatalf(err.Error())
	}
	err = os.Setenv("DB_HOST", env["DB_HOST"])
	err = os.Setenv("DB_PORT", env["DB_PORT"])
	err = os.Setenv("DB_USER", env["DB_USER"])
	err = os.Setenv("DB_NAME", env["DB_NAME"])
	err = os.Setenv("DB_PASSWORD", env["DB_PASSWORD"])
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
