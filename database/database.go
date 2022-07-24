package database

import (
	"fmt"
	"log"
	"os"

	"github.com/onattech/go-postgres-docker-heroku/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// connectDb
func ConnectDb() {
	var dsn string

	if os.Getenv("ENV") == "PROD" {
		dsn = os.Getenv("DATABASE_URL")
	} else {
		dsn = fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v port=%v",
			os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"),
			os.Getenv("PG_DBNM"), os.Getenv("PG_PORT"),
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Product{})

	DB = db

}
