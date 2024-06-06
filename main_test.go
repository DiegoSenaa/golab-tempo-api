package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandleRequest(t *testing.T) {
	app := fiber.New()
	app.Get("/:cep", handleRequest)

	// Test for invalid zipcode length
	req := httptest.NewRequest("GET", "/1234567", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)

	// Mock fetchData function
	originalFetchData := fetchDataFunc
	fetchDataFunc = func(url string) ([]byte, error) {
		if url == urlViaCep+"ws/12345678/json" {
			return json.Marshal(ViaCepResponse{
				Localidade: "Sao Paulo",
				Uf:         "SP",
			})
		}
		if strings.Contains(url, "weatherapi.com") {
			return json.Marshal(WeatherApiResponse{
				Current: struct {
					LastUpdatedEpoch int     `json:"last_updated_epoch"`
					LastUpdated      string  `json:"last_updated"`
					TempC            float64 `json:"temp_c"`
					TempF            float64 `json:"temp_f"`
					IsDay            int     `json:"is_day"`
					Condition        struct {
						Text string `json:"text"`
						Icon string `json:"icon"`
						Code int    `json:"code"`
					} `json:"condition"`
					WindMph    float64 `json:"wind_mph"`
					WindKph    float64 `json:"wind_kph"`
					WindDegree int     `json:"wind_degree"`
					WindDir    string  `json:"wind_dir"`
					PressureMb float64 `json:"pressure_mb"`
					PressureIn float64 `json:"pressure_in"`
					PrecipMm   float64 `json:"precip_mm"`
					PrecipIn   float64 `json:"precip_in"`
					Humidity   int     `json:"humidity"`
					Cloud      int     `json:"cloud"`
					FeelslikeC float64 `json:"feelslike_c"`
					FeelslikeF float64 `json:"feelslike_f"`
					VisKm      float64 `json:"vis_km"`
					VisMiles   float64 `json:"vis_miles"`
					Uv         float64 `json:"uv"`
					GustMph    float64 `json:"gust_mph"`
					GustKph    float64 `json:"gust_kph"`
				}{TempC: 25.0, TempF: 77.0, GustKph: 15.0},
			})
		}
		return nil, errors.New("unknown URL")
	}
	defer func() { fetchDataFunc = originalFetchData }()

	// Test for valid zipcode
	req = httptest.NewRequest("GET", "/12345678", nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, 25.0, body["temp_C"])
	assert.Equal(t, 77.0, body["temp_F"])
	assert.Equal(t, 298.15, body["temp_K"])
}

func TestFetchData(t *testing.T) {
	// Test successful data fetch
	url := "https://viacep.com.br/ws/01001000/json"
	_, err := fetchData(url)
	assert.Nil(t, err)

	// Test with an invalid URL
	_, err = fetchData("https://invalid.url")
	assert.NotNil(t, err)
}

func TestRemoveAccents(t *testing.T) {
	assert.Equal(t, "Sao Paulo", removeAccents("São Paulo"))
}
