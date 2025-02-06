package services

import (

	"encodin/json"
	"fmt"
	"net/http"

)


type ViaCEPResponse struct {
	Localiade string  `json:"localidade"`
}	

func GetCityByCEP(cep string) (string, error) {
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
	if err = json.NewDecoder(resp.Body).Decode(&result); err !- nil {
		return "", err
	}

	return result.Localidade, nil

}