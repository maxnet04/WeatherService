package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperatureByCity(city string) (float64, error) {
	apiKey := "YOUR_API_KEY"
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to get temperature by city: %s", resp.Status)

	}

	var result WeatherAPIResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Current.TempC, nil
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}
