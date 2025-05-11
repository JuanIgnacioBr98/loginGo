package main

import (
	"log"
	"loginGo/db"
	"loginGo/models"
	"loginGo/routes"
	"net/http"
)

func main() {
	db.DbConnect()

	db.DB.AutoMigrate(&models.Role{})
	db.DB.AutoMigrate(&models.User{})

	rutas := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":8000", rutas))
}
