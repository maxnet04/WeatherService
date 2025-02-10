package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/maxnet04/WeatherService/services"
	"github.com/stretchr/testify/assert"
)

type MockWeatherService struct{}

func (m *MockWeatherService) GetCityByCEP(cep string) (string, error) {
	return "São Paulo", nil
}

func (m *MockWeatherService) GetTemperatureByCity(city string) (float64, error) {
	return 20, nil
}

func (m *MockWeatherService) CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func (m *MockWeatherService) CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func TestGetWeatherSucess(t *testing.T) {

	//Cria um novo request com um CEP valido
	mockService := &MockWeatherService{}
	req, err := http.NewRequest("GET", "/weather/12345678", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Usar mux para definir variáveis de rota
	req = mux.SetURLVars(req, map[string]string{"cep": "12345678"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWeather(mockService))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response WeatherResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 20.0, response.TempC)
	assert.Equal(t, 68.0, response.TempF)
	assert.Equal(t, 293.15, response.TempK)
}

// TestGetWeatherInvalidCeo testa o handler GetWeather com um CEP invalido
func TestGetWeatherInvalidCeo(t *testing.T) {

	//Cria um novo request com um CEP valido
	req, err := http.NewRequest("GET", "/weather/12345", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Usar mux para definir variáveis de rota
	req = mux.SetURLVars(req, map[string]string{"cep": "12345"})

	service := services.WeatherService(&services.RealWeatherService{})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWeather(service))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

// TestGetWeatherCEPNotFound testa o handler GetWeather com um CEP que não foi encontrado
func TestGetWeatherCEPNotFound(t *testing.T) {

	//Cria um novo request com um CEP valido
	req, err := http.NewRequest("GET", "/weather/12345678s", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Usar mux para definir variáveis de rota
	req = mux.SetURLVars(req, map[string]string{"cep": "12345678"})

	service := services.WeatherService(&services.RealWeatherService{})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWeather(service))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
