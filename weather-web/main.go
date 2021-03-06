package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type AutoGenerated struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

type Show struct {
	City        string  `json:"city"`
	Description string  `json:"description"`
	Temp        float64 `json:"temp"`
}

func CityWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["city"]

	city := []string{name}

	if name == "all" {
		city = []string{"bangkok", "newyork", "hobart", "nairobi", "kupang"}
	}

	var show []Show
	//show := make([]Show, 1)

	for i, cityName := range city {

		fmt.Println(i, cityName)

		url := "http://127.0.0.1:8882/api/v1/weather/" + cityName
		res, _ := http.Get(url)

		var web AutoGenerated
		err := json.NewDecoder(res.Body).Decode(&web)
		if err != nil {
			log.Fatal(err)
		}

		showArr := Show{City: web.Name, Description: web.Weather[0].Description, Temp: web.Main.Temp}
		show = append(show, showArr)
	}

	json.NewEncoder(w).Encode(show)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("Start at ", start)
		next.ServeHTTP(w, r)
		fmt.Println("Completed in ", time.Since(start))
	})
}
func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/weather/{city}", CityWeather).Methods("GET")
	r.Use(loggingMiddleware)

	return r
}
func main() {
	http.ListenAndServe(":8080", NewRouter())
}
