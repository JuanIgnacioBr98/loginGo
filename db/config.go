package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var CONNECTION_STRING = "host=localhost user=postgres password=postgres dbname=logingo port=5432 sslmode=disable TimeZone=America/Argentina/Buenos_Aires"
var DB *gorm.DB

func DbConnect() {
	var err error
	DB, err = gorm.Open(postgres.Open(CONNECTION_STRING))

	if err != nil {
		log.Fatal("Error de conexión a la base de datos: ", err)
		return
	}

	log.Println("Conección exitosa a la base de datos")
}
