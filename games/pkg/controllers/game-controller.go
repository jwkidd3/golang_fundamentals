package controllers

import (
	"encoding/json"
	"github.com/cyberdr0id/go-rest-api/pkg/models"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var games []models.Game // for init

func UpdateGame(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	for index, item := range models.Games {
		if item.ID == params["id"] {
			models.Games = append(models.Games[:index], models.Games[index+1:]...)
			var movie models.Game
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["id"]
			models.Games = append(models.Games, movie)
			err := json.NewEncoder(writer).Encode(models.Games)
			if err != nil {
				log.Fatalf("error with encoding updating game: %s", err.Error())
				return
			}

			return
		}
	}
}

func CreateGame(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var game models.Game

	_ = json.NewDecoder(request.Body).Decode(&game)

	game.ID = strconv.Itoa(rand.Intn(1000000))
	models.Games = append(models.Games, game)

	err := json.NewEncoder(writer).Encode(models.Games)
	if err != nil {
		log.Fatalf("error with encoding creating game: %s", err.Error())
		return
	}
}

func GetGameById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for _, item := range models.Games {
		if item.ID == params["id"] {
			err := json.NewEncoder(writer).Encode(item)
			if err != nil {
				log.Fatalf("error with encoding geting game: %s", err.Error())
				return
			}

			return
		}
	}
}

func DeleteGame(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for index, item := range models.Games {
		if item.ID == params["id"] {
			models.Games = append(models.Games[:index], models.Games[index+1:]...)
		}
	}

	err := json.NewEncoder(writer).Encode(models.Games)
	if err != nil {
		log.Fatalf("error with encoding deleting game: %s", err.Error())
		return
	}
}

func GetGames(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(models.Games)

	if err != nil {
		log.Fatalf("error with encoding geting games: %s", err.Error())
		return
	}
}
