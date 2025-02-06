package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Converting temperatures")
	var temperatureType string
	fmt.Println("Please enter the type of temperature you wish to convert to")
	fmt.Scanln(&temperatureType)

	if temperatureType != "c" && temperatureType != "f" && !strings.EqualFold("celsius", temperatureType) && !strings.EqualFold("fahrenheit", temperatureType) {
		fmt.Println("Unknown temperature type provided, exiting")
		os.Exit(1)
	} else if temperatureType == "c" || temperatureType == "f" {
		switch temperatureType {
		case "c":
			fmt.Println("Celsius has been selected")
		case "f":
			fmt.Println("Fahrenheit has been selected")
		}
	} else {
		fmt.Println(temperatureType + " has been selected")
	}
	fmt.Println("Fahrenheit is worked out by taking Celsius and multipling by 1.8, then adding 32.")

	celsius := 10
	fmt.Printf("Celsius is: %d\n", celsius)
	farenheit := (float64(celsius) * 1.8) + 32

	fmt.Printf("Farenheit is: %gf based on Celsisus being: %d\n ", farenheit, celsius)

	farenheit = 60
	fmt.Printf("Farenheit is: %g\n", farenheit)

	celsius = int((farenheit - 32) / 1.8)

	fmt.Printf("Celsius is: %d where farenheit is: %g\n", celsius, farenheit)

}
