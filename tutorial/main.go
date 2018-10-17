package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func query(city string) (weatherData, error) {

}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		TempKelvin float64 `json:"temp"`
	} `json:"main"`
}

const (
	// https://openweathermap.org/ more info
	weatherAPIEndpoint string = "http://api.openweathermap.org/data/2.5/weather"
	weatherAPIKey      string = "cd6d8b4025c79074ebaec745f80919df"
)
