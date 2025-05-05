package routes

import (
	"loginGo/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()

	api.HandleFunc("", controllers.InitRoute).Methods("GET")
	api.HandleFunc("/", controllers.InitRoute).Methods("GET")

	return rutas
}
