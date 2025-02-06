package main

import (
	"log"
	"net/http"

	"github.com/gorrola/mux"
	"github.com/maxnet04/weatherservice/handlers"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", handlers.GetWeather).Methods("GET")

	log.Println("Server started on: port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
