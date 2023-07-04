package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func connect() {
	dsn := "host=localhost user=postgres password=pwd123 dbname=gorm port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open database!\n", err)
		os.Exit(2)
	}

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// TODO: make migrations
	Database = DbInstance{Db: db}
}
