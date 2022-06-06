package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("OPENWEATHERMAP_API_KEY")
var lat = os.Getenv("LAT")
var long = os.Getenv("LONG")

func main() {

	w, err := owm.NewCurrent("F", "EN", apiKey) // (internal - OpenWeatherMap reference for Farenheit) with English output
	lon, err := strconv.ParseFloat(long, 64)
	latitude, err := strconv.ParseFloat(lat, 64)
	fmt.Println(lon)
	fmt.Println(latitude)
	if err != nil {
		log.Fatalln(err)
	}
	w.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: lon,
			Latitude:  latitude,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
	owm.ValidAPIKey(apiKey)

	fmt.Println(w)

}
