package main

import (
	"github.com/gocrud/intializer"
	"github.com/gocrud/models"
)

func init() {
	intializer.Databasecon()
}
func main() {

	// create schema
	intializer.DB.AutoMigrate(&models.User{})

}
