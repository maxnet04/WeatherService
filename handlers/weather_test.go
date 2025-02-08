package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// TestGetWeatherInvalidCeo testa o handler GetWeather com um CEP invalido
func TestGetWeatherInvalidCeo(t *testing.T) {

	//Cria um novo request com um CEP valido
	req, err := http.NewRequest("GET", "/weather/12345", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Usar mux para definir variáveis de rota
	req = mux.SetURLVars(req, map[string]string{"cep": "12345"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWeather)
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

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWeather)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
