package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/maxnet04/WeatherService/handlers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", handlers.GetWeather).Methods("GET")

	log.Println("Server started on: port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
