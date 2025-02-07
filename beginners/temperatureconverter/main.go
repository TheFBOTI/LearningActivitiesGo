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

	if !strings.EqualFold("c", temperatureType) && !strings.EqualFold("f", temperatureType) && !strings.EqualFold("celsius", temperatureType) && !strings.EqualFold("fahrenheit", temperatureType) {
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

	var temperatureValue int
	fmt.Println("Please enter the value of temperature you want to be converted")
	fmt.Scanln(&temperatureValue)

	var celsius int
	var fahrenheit float64
	if strings.EqualFold("c", temperatureType) || strings.EqualFold("celsius", temperatureType) {
		fmt.Println("Fahrenheit is worked out by taking Celsius and multiplying by 1.8, then adding 32.")
		celsius = temperatureValue
		fmt.Printf("Celsius is: %d\n", celsius)
		fahrenheit = (float64(celsius) * 1.8) + 32

		fmt.Printf("Farenheit is: %gf based on Celsisus being: %d\n ", fahrenheit, celsius)
	} else {
		fmt.Println("Celsius is worked out by taking Fahrenheit subtracting 32 and then dividing by 1.8.")
		fahrenheit = float64(temperatureValue)
		fmt.Printf("Fahrenheit is: %g\n", fahrenheit)
		celsius = int((fahrenheit - 32) / 1.8)
		fmt.Printf("Celsius is: %d where fahrenheit is: %g\n", celsius, fahrenheit)
	}
}
