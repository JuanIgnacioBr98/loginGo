package routes

import (
	"loginGo/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()
	api.HandleFunc("", controllers.InitRoute).Methods("GET")

	roles := api.PathPrefix("/roles").Subrouter()
	roles.HandleFunc("", controllers.GetRoles).Methods("GET")
	roles.HandleFunc("/{id}", controllers.GetOneRole).Methods("GET")
	roles.HandleFunc("", controllers.NewRole).Methods("POST")
	roles.HandleFunc("/{id}", controllers.DeleteRole).Methods("DELETE")

	return rutas
}
