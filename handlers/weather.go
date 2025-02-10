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

func GetWeather(service services.WeatherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		cep := vars["cep"]
		if len(cep) != 8 {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		city, err := service.GetCityByCEP(cep)
		if err != nil {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}

		tempC, err := service.GetTemperatureByCity(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := WeatherResponse{
			TempC: tempC,
			TempF: service.CelsiusToFahrenheit(tempC),
			TempK: service.CelsiusToKelvin(tempC),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
