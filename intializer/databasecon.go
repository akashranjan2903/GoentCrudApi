package intializer

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Databasecon() {

	var err error
	dsn := "host=localhost user=postgres password=@kash123 dbname=DemoDatabase port=8080 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		log.Fatal("Failed to connect to Database")
	}

}
