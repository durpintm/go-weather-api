package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type WeatherRequest struct {
	Name string `json:"name"`
}

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

const Dport = ":8000"

func main() {

	// router := mux.NewRouter()

	// router.HandleFunc("/weather/city", getWeatherByCityHandler).Methods("GET")
	// router.HandleFunc("/weather/city", weatherHandler).Methods("POST")

	// fmt.Printf("Server is starting on port: %v\n", Dport)

	// http.ListenAndServe(Dport, router)

	apiKey, err := loadApiConfig(".apiConfig")
	if err != nil {
		fmt.Println("Error accessing api key", err)
	}
	fmt.Println(apiKey.OpenWeatherMapApiKey)
}

// This function gets the api key from the config file
func loadApiConfig(filename string) (*apiConfigData, error) {
	// Open the api config file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return nil, err
	}
	defer file.Close()

	// Decode the api config file
	var config apiConfigData
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return nil, err
	}
	return &config, nil
}

func getWeatherByCityHandler(w http.ResponseWriter, r *http.Request) {

    apiKey,err := loadApiConfig(".apiConfig")
    if(err != nil){
        fmt.Println("Error loading api config data", err)
    }
	// Get the city from the query parameters
	cityName := r.URL.Query().Get("name")
	units := "metric" // For temperature unit in Celcius

	// Constructing the URL for the API request
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&appid=%s", cityName, units, apiKey.OpenWeatherMapApiKey)

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
     var weatherResp WeatherResponse
    if err := json.NewDecoder(response.Body).Decode(&weatherResp); err != nil {
        fmt.Printf("Error decoding JSON: %v\n", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherResp)
}

}
// This api handler accepts a JSON body with a city name and returns the current weather in JSON format
func weatherHandler(w http.ResponseWriter, r *http.Request) {

    apiKey,err := loadApiConfig(".apiConfig")
    if(err != nil){
        fmt.Println("Error loading api config data", err)
    }

    var weatherReq WeatherRequest
    if err := json.NewDecoder(r.Body).Decode(&weatherReq); err != nil {
        fmt.Printf("Error decoding JSON: %v\n", err)
        return
    }

	// Get the city from the query parameters
	cityName := weatherReq.Name
	units := "metric" // For temperature unit in Celcius

	// Constructing the URL for the API request
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&appid=%s", cityName, units, apiKey.OpenWeatherMapApiKey)

	response, err := http.Get(url)
	
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
     var weatherResp WeatherResponse
    if err := json.NewDecoder(response.Body).Decode(&weatherResp); err != nil {
        fmt.Printf("Error decoding JSON: %v\n", err)
        return
    }

    fmt.Println(weatherResp)
    fmt.Println(response)
    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherResp)
}
