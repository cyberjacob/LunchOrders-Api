package main

import (
	"github.com/cuvva/ksuid"
	"github.com/cyberjacob/LunchOrders-Api/api"
	"github.com/cyberjacob/LunchOrders-Api/models/db"
	"log"
)

var Db *gorm.DB

// our main function
func main() {
	log.Printf("Server started")
	ksuid.SetEnvironment("dev")

	db, err := db.MakeDB()
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.OrderItems{})
	db.AutoMigrate(&models.Session{})

	router := api.NewRouter()

	router.Run(":8456")
}
