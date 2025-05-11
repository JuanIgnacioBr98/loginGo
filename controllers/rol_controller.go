package controllers

import (
	"encoding/json"
	"loginGo/db"
	"loginGo/models"
	"loginGo/utils"
	"net/http"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	var roles []models.Role
	db.DB.Find(&roles)
	response := utils.Response{
		Msg:       "Lista de Roles",
		Data:      "",
		SatusCode: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetOneRole(w http.ResponseWriter, r *http.Request) {
	response := utils.Response{
		Msg:       "Buscar",
		Data:      "Buscar uno",
		SatusCode: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func NewRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var role models.Role

	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		response := utils.Response{
			Msg:       "Error en los datos enviados",
			Data:      err.Error(),
			SatusCode: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	newRole := db.DB.Create(&role)

	err := newRole.Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := utils.Response{
			Msg:       "Error al crear el rol",
			Data:      newRole.Error.Error(),
			SatusCode: http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	response := utils.Response{
		Msg:       "Rol creado con éxito",
		Data:      role,
		SatusCode: http.StatusCreated,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	response := utils.Response{
		Msg:       "Eliminar Rol",
		Data:      "Eliminado",
		SatusCode: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}
