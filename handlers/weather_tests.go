package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TestGetWeatherSucess testa o handler GetWeather com um CEP valido
func TestGetWeatherSucess(t *testing.T) {

	//Cria um novo request com um CEP valido
	req, err := http.NewRequest("GET", "/weather/01001000", nil)
	if err != nil {
		t.Fatal(err)
	}

	//Cria um ResponseRecorder para capturar a rresposta da requisição
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", GetWeather).Methods("GET")
	r.ServeHTTP(rr, req)

	//Verifica se o status code é 200
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code incorreto: Retornado: %v Esperado: %v",
			status, http.StatusOK)
	}
}

// TestGetWeatherInvalidCeo testa o handler GetWeather com um CEP invalido
func TestGetWeatherInvalidCeo(t *testing.T) {

	req, err := http.NewRequest("GET", "/weather/012", nil) // Cep invalido
	if err != nil {
		t.Fatal(err)
	}

	//Cria um ResponseRecorder para capturar a rresposta da requisição
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", GetWeather).Methods("GET")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Status code incorreto: Retornado: %v Esperado: %v",
			status, http.StatusUnprocessableEntity)
	}

	expected := "CEP must have 8 digits\n"
	if rr.Body.String() != expected {
		if rr.Body.String() != expected {
			t.Errorf("Body incorreto: Retornado: %v Esperado: %v",
				rr.Body.String(), expected)
		}
	}
}

// TestGetWeatherCEPNotFound testa o handler GetWeather com um CEP que não foi encontrado
func TestGetWeatherCEPNotFound(t *testing.T) {

	req, err := http.NewRequest("GET", "/weather/88888888", nil) // Cep não existe
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", GetWeather).Methods("GET")
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Status code incorreto: Retornado: %v Esperado: %v",
			status, http.StatusNotFound)
	}

	expected := "failed to get city by CEP: 404 Not Found\n"
	if rr.Body.String() != expected {
		t.Errorf("Body incorreto: Retornado: %v Esperado: %v",
			rr.Body.String(), expected)
	}
}
