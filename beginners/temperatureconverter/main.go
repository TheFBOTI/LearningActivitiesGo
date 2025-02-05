package main

import (
	"fmt"
)

func main() {
	fmt.Println("Converting temperatures")

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
