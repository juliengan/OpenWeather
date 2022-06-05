package main

import (
	"fmt"
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("OPENWEATHERMAP_API_KEY")

func main() {

	w, err := owm.NewCurrent("F", "ru", apiKey) // fahrenheit (imperial) with Russian output
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Phoenix")
	fmt.Println(w)
}
