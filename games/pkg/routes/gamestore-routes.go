package routes

import (
	"github.com/cyberdr0id/go-rest-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterGameStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/game", controllers.CreateGame).Methods("POST")
	router.HandleFunc("/game", controllers.GetGames).Methods("GET")
	router.HandleFunc("/game/{id}", controllers.GetGameById).Methods("GET")
	router.HandleFunc("/game/{id}", controllers.UpdateGame).Methods("PUT")
	router.HandleFunc("/game/{id}", controllers.DeleteGame).Methods("DELETE")

}
