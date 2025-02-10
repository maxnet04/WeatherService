package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherService interface {
	GetCityByCEP(cep string) (string, error)
	GetTemperatureByCity(city string) (float64, error)
	CelsiusToFahrenheit(celsius float64) float64
	CelsiusToKelvin(celcius float64) float64
}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro,omitempty"`
}

type RealWeatherService struct{}

func (s *RealWeatherService) GetCityByCEP(cep string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get city by CEP: %s", resp.Status)
	}

	var result ViaCEPResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.Erro {
		return "", fmt.Errorf("can not find zipcode")
	}

	return result.Localidade, nil
}

func (s *RealWeatherService) GetTemperatureByCity(city string) (float64, error) {
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

func (s *RealWeatherService) CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func (s *RealWeatherService) CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}
