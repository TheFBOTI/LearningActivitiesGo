package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fileToSearch string
	fmt.Print("Please Enter the Directory and the file name you would like to search: \n")
	fmt.Scanln(&fileToSearch)
	fmt.Println("You have typed in " + fileToSearch + " ---- Is this correct? Y/N?")
	var isThisTheFile string
	fmt.Scanln(&isThisTheFile)
	isThisTheFile = strings.ToLower(isThisTheFile)
	switch isThisTheFile {
	case "yes":
		fmt.Println("Continuing")
	case "y":
		fmt.Println("Continuing")
	case "no":
		fmt.Println("Not the file expected, restarting process")
		main()
	case "n":
		fmt.Println("Not the file expected, restarting process")
		main()
	default:
		fmt.Println("Unexpected input, restarting; yes, no, y or n are supported values. \n Thanks Fboti ~ ")
		main()
	}

	searchForWord(fileToSearch)

}

func searchForWord(fileToSearch string) {
	var searchWord string
	fmt.Print("Enter the word you'd like to search for: ")
	fmt.Scanln(&searchWord)

	countWordFrequency(fileToSearch, strings.ToLower(searchWord))
	var retry string
	fmt.Print("Would you like to search again in the same file? yes or y to redo")
	fmt.Scanln(&retry)
	retry = strings.ToLower(retry)
	switch retry {
	case "yes":
		searchForWord(fileToSearch)
	case "y":
		searchForWord(fileToSearch)
	default:
		fmt.Println("Ending, thanks for trying the tool ~")
	}

}

func countWordFrequency(filename string, wordToSearchFor string) {
	fmt.Println("Searching: " + filename + " for " + wordToSearchFor)
	dataFromFile, err := os.ReadFile(filename)
	wordToSearchForSanitised := strings.ToLower(strings.Trim(wordToSearchFor, ".,!?:;"))

	if err != nil {
		log.Fatal(err)
	}
	splitText := strings.Split(string(dataFromFile), " ")
	for _, word := range splitText {
		cleanedWordFromFile := strings.ToLower(strings.Trim(word, ".,!?:;"))
		if cleanedWordFromFile == wordToSearchForSanitised {
			fmt.Println("The word '"+wordToSearchForSanitised+"' appears", countOccurrences(splitText, wordToSearchForSanitised), "times.")
			return
		}
	}
	fmt.Println("Sorry, the word '" + wordToSearchFor + "' was not found in the text.")
}

func countOccurrences(dataFromFile []string, wordToSearchFor string) int {
	count := 0
	for _, word := range dataFromFile {
		if strings.EqualFold(strings.Trim(word, ".,!?:;"), wordToSearchFor) {
			count++
		}
	}
	return count
}
