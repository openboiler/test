package main

import (
	"fmt"
	"log"
	"os"

	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
)

func main() {
	os.Setenv("OWM_API_KEY", "21ed659897b27e258c206a878fc6e5db")
	w, err := owm.NewCurrent("C", "it") // fahrenheit (imperial) with Russian output
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Padua")
	fmt.Printf("%+v\n", w)
}
