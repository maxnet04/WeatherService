package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/maxnet04/WeatherService/handlers"
	"github.com/maxnet04/WeatherService/services"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	watherService := services.WeatherService(&services.RealWeatherService{})

	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", handlers.GetWeather(watherService)).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Weather Service")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
	log.Printf("Server started on: port :%s", port)

}
