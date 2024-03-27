package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type apiConfigData struct{
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type WeatherRequest struct{
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

func main(){

    // router := mux.NewRouter()

    // router.HandleFunc("/weather/city", getWeatherByCityHandler).Methods("GET")
    // router.HandleFunc("/weather/city", weatherHandler).Methods("POST")

	// fmt.Printf("Server is starting on port: %v\n", Dport)

	// http.ListenAndServe(Dport, router)

    apiKey,err := loadApiConfig(".apiConfig")
    if(err != nil){
        fmt.Println("Error accessing api key", err)
    }
    fmt.Println(apiKey.OpenWeatherMapApiKey)
}

// This function gets the api key from the config file
func loadApiConfig(filename string) (*apiConfigData, error){
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
