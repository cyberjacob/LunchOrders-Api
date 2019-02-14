package main

import (
	"github.com/cuvva/ksuid"
	"github.com/cyberjacob/LunchOrders-Api/api"
	"github.com/cyberjacob/LunchOrders-Api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var Db *gorm.DB

// our main function
func main() {
	log.Printf("Server started")
	ksuid.SetEnvironment("dev")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.OrderItems{})
	db.AutoMigrate(&models.Session{})

	models.Db = db

	router := api.NewRouter()

	router.Run(":8456")
}
