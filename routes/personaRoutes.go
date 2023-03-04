package routes

import (
	"api-go-4softwaredeveloper/controllers"

	"github.com/gorilla/mux"
)

func SetPersonasRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/persona/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
}
