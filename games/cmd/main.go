package main

import (
	"fmt"
	"github.com/cyberdr0id/go-rest-api/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterGameStoreRoutes(r)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}