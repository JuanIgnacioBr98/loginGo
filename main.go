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
	rutas := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":8000", rutas))
}
