# Weather Api in Go

## How to build docker image and docker file and how to push it to docker hub:

```bash
docker build -t go-weather-api .
docker run -p 8000:8000 -t go-weather-api
docker tag go-weather-api:latest durpintm/go-weather-api-hub:latest
docker push durpintm/go-weather-api-hub:latest

docker run -p 8000:8000 durpintm/go-weather-api-hub:latest
```

## Test the API

```bash
### City Weather using GET request
curl -X GET http://localhost:8000/weather/city?name=toronto
### Post an item to the list
curl -X POST http://localhost:8000/weather/city -H "Content-Type: application/json" -d "{\"name\":\"Sydney\"}"
```

## Overview:

- The Key objective of this project is to create an API server with the help of Go programming language to acknowledge GET and POST requests with details about the current weather in a particular city. The main outcome of this assignment is to create an API server in Go, holding API integrations, containerizing the applications to be run using docker, and significantly managing code using version control.

Set Up and Run the Application:

1. Setup Development Environment: Install Go and Docker.
2. Create Go Project: Initialize a new Go module and define project structure.
3. Define Data Model: Create Go structs to represent weather data and API responses.
4. HTTP Routing: Use a Go web framework (e.g., Gorilla Mux) for handling HTTP routing.
5. Integrate Weather API(OpenWeatherMap.org): Implement functions to fetch weather data from the chosen API.
6. Handle GET Requests: Create handlers to respond to GET requests for weather information.
7. Handle POST Requests: Implement handlers for POST requests allowing users to specify a city.
8. Return JSON Responses: Serialize weather data into JSON format for response.
9. Containerize Application: Write a Dockerfile to package the Go application into a container.
10. Build and Run Docker Container: Build the Docker image and run a containerized instance of the API server.
11. Test API Endpoints: Use tools like cURL or Postman to test the functionality of the API endpoints.
12. Deploy Containerized Application: Deploy the Docker container to the docker hub.

### Note: Update API key in .apiConfig file

## About External Weather API:

- In this Project we have used a external weather API, named as OpenWeatherMap. Basically, this API provides weather data in terms of current weather, historical data, or forecasts at many locations around the world.By the help of this API we can get weather data such as, humidity, Air Pressure, UV, visiblity and many more.Moreover, this API comprises geocoding such as latitude and longitude and allowing users to search for weather information of a particular locations. For this, Users can sign up for an API key to access the OpenWeatherMap API. for the purpose of authentication requests and track usage API key is required.
