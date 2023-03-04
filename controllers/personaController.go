package controllers

import (
	"api-go-4softwaredeveloper/commons"
	"api-go-4softwaredeveloper/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConexion()
	defer db.Close()

	db.Find(&personas)
	json, _ := json.Marshal(personas)
	commons.SendResponse(w, http.StatusOK, json)
}

func Save(w http.ResponseWriter, r *http.Request) {
	persona := models.Persona{}
	db := commons.GetConexion()
	defer db.Close()

	error := json.NewDecoder(r.Body).Decode(&persona)

	if error != nil {
		log.Fatal(error)
		commons.SendError(w, http.StatusBadRequest)
		return
	}

	error = db.Save(&persona).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(w, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)

	commons.SendResponse(w, http.StatusCreated, json)
}
