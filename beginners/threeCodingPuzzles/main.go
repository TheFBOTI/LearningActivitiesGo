package main

import (
	"fmt"
	"strings"
)

func main() {
	var puzzleChoice string
	fmt.Print("Please select what puzzle you would like: \n Fibonacci \n Palindrome checker \n Prime Number Generator")
	fmt.Scan(&puzzleChoice)

	switch strings.ToLower(puzzleChoice) {
	case "fibonacci":
		//function call

	case "palindrome checker":
		//function call

	case "prime number":
		//function call

	}

	var restart string
	fmt.Println("Would you like to restart?")
	fmt.Scan(&restart)
	if strings.ToLower(restart) == "yes" || strings.ToLower(restart) == "y" {
		main()
	}
}
