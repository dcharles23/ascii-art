package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		userInput := os.Args[1]
		fileLines := ReadStandardTxt()

		asciiTemplates := returnToTheASCIIArray(fileLines)

		printOutput(userInput, asciiTemplates)
		//for each letter of the user input it asigns an ascii character from asciiTemplates
	}
}

func ReadStandardTxt() []string {
	readFile, err := os.Open("standard.txt") // os.Open opens the file so data can be accessed
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}

// for every 8 lines a new ascii symbol is added
func returnToTheASCIIArray(fileLines []string) [][]string {
	var asciiTemplates [][]string // 2D array is an array that stores an array
	counter := 0
	var tempAsciArray []string
	for _, line := range fileLines {
		counter++
		// fmt.Println(index, line)
		if counter != 1 {
			tempAsciArray = append(tempAsciArray, line) //tempAsciArray is the current symbol
		}
		if counter == 9 {
			asciiTemplates = append(asciiTemplates, tempAsciArray)
			counter = 0         // resets counter
			tempAsciArray = nil // resets tempAsciArray
		}
	}
	return asciiTemplates
}

// problem '\n' logic when we detect 2 '\n' or 1 '\n' is different
func printMultipleCharacter(s string, asciiTemplates [][]string) {
	// for example 'hello'
	// index 0 follwed by a $ and newline, index 1 follwed by a $ and newline etc.
	// the $ decalres the end of terminal usage
	tempIntArrLetter := returnAsciiCodeInt(s)
	for i := 0; i < 8; i++ { //each ascii symbol is made up of 8 lines
		for _, v := range tempIntArrLetter {
			fmt.Print(asciiTemplates[v][i])
		}
		fmt.Println("$")
	}
}

func returnAsciiCodeInt(s string) []int {
	var tempIntArrLetter []int
	for _, v := range s {
		tempIntArrLetter = append(tempIntArrLetter, (int(v) - 32)) // 32 is the decimal number position for space in the ascii table. in the standard txt file we store a space as the 0 index
	}
	return tempIntArrLetter
}

func printOutput(text string, asciiTemplates [][]string) {
	substrings := returnstring2EndlineArray(text)
	lenOfsubstrings := len(substrings)
	for index, v := range substrings {
		if v == "\\n" {
			// If it is last one
			if index == lenOfsubstrings-1 {
				fmt.Println("$")
			} else if index == 0 {
				fmt.Println("$")
			} else {
				if substrings[index-1] == "\\n" { // if you want to detect a '\n' in a string you need '\\n' we need to detect the actual '\n' not just a new line
					fmt.Println("$")

				} else {
					// "Hello\nWorld"
				}
			}
		} else {
			printMultipleCharacter(v, asciiTemplates)
		}
	}

}

func returnstring2EndlineArray(text string) []string {
	substrings := make([]string, 0)
	escapedN := "\\n"
	newline := "\n"

	for {
		index := strings.Index(text, escapedN)
		if index == -1 {
			substrings = append(substrings, text)
			break
		}

		substrings = append(substrings, text[:index])

		if index+len(escapedN) < len(text) && text[index+len(escapedN)] == 'n' {
			substrings = append(substrings, newline)
			text = text[index+len(escapedN)+1:]
		} else {
			substrings = append(substrings, escapedN)
			text = text[index+len(escapedN):]
		}
	}
	// fmt.Printf("%#v\n", substrings)
	var mysubstring2 []string
	for _, mysub := range substrings {
		if mysub != "" {
			mysubstring2 = append(mysubstring2, mysub)
		}
	}
	// fmt.Printf("%#v\n", mysubstring2)
	return mysubstring2
}
