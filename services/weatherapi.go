package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperatureByCity(city string) (float64, error) {
	apiKey := os.Getenv("WEATHERAPI_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("WEATHERAPI_KEY not set")
	}

	req, err := http.NewRequest("GET", "http://api.weatherapi.com/v1/current.json", nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	q.Add("key", apiKey)
	q.Add("q", city)
	q.Add("aqi", "no")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result WeatherAPIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
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
