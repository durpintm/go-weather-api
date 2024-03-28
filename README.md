# Weather Api
Overview:
The Key objective of this project is to create an API server with the help of  Go programming language to acknowlede GET and POST requests with details about the current weather in a particular city. The main outcome of Cerating  a sturdy  API server in Go, holding API integrations, deploying containerized applications, and significantly managing code using version control.

Set Up and Run the Application:
1.Setup Development Environment: Install Go and Docker.
2.Select Weather API: In this we selcted  an external weather API provider i.e, OpenWeatherMap.
3.Create Go Project: Initialize a new Go module and define project structure.
4.Define Data Model: Create Go structs to represent weather data and API responses.
5.HTTP Routing: Use a Go web framework (e.g., Gorilla Mux) for handling HTTP routing.
6.Integrate Weather API(OpenWeatherMap.org): Implement functions to fetch weather data from the chosen API.
7.Handle GET Requests: Create handlers to respond to GET requests for weather information.
8.Handle POST Requests: Implement handlers for POST requests allowing users to specify a city.
9.Return JSON Responses: Serialize weather data into JSON format for response.
10.Containerize Application: Write a Dockerfile to package the Go application into a container.
11.Build and Run Docker Container: Build the Docker image and run a containerized instance of the API server.
12.Test API Endpoints: Use tools like cURL or Postman to test the functionality of the API endpoints.
13.Deploy Containerized Application: Deploy the Docker container to a hosting service for public access.

About External Weather API:
In this Project we have used a external weather API, named as OpenWeatherMap. Basically, this API provides weather data in terms of current weather, historical data, or forecasts at many locations around the world.By the help of this API we can get weather data such as, humidity, Air Pressure, UV, visiblity and many more.Moreover, this API comprises geocoding such as latitude and longitude and  allowing users to search for weather information of a particular locations. For this, Users can sign up for an API key to access the OpenWeatherMap API.  for the purpose of authentication requests and track usage API key is required.





 