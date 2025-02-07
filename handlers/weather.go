package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxnet04/WeatherService/services"
)

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cep := vars["cep"]
	if len(cep) != 8 {
		http.Error(w, "CEP must have 8 digits", http.StatusUnprocessableEntity)
		return
	}

	city, err := services.GetCityByCEP(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tempC, err := services.GetTemperatureByCity(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := WeatherResponse{
		TempC: tempC,
		TempF: services.CelsiusToFahrenheit(tempC),
		TempK: services.CelsiusToKelvin(tempC),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
