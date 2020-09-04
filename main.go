package main

import (
	"gorm-demo/database"
	"gorm-demo/models"

	//"gorm-demo/models"
)

func main() {
	//person := &models.Person{Name: "shubham"}
	db := database.ConnectToPostgres()


	contact := models.Contact{MobileNumber: "892161853", EmailId:"developer.shubhamv@gmail.com"}
	// Create
	db.Create(&models.Person{Name: "shubhamv", Contact: contact})
}
