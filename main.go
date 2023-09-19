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

	}
}

func ReadStandardTxt() []string {
	readFile, err := os.Open("standard.txt")
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

func returnToTheASCIIArray(fileLines []string) [][]string {
	var asciiTemplates [][]string
	counter := 0
	var tempAsciArray []string
	for _, line := range fileLines {
		counter++
		// fmt.Println(index, line)
		if counter != 1 {
			tempAsciArray = append(tempAsciArray, line)
		}
		if counter == 9 {
			// fmt.Println("add to list") // but dont include the first line because it is empty line
			asciiTemplates = append(asciiTemplates, tempAsciArray)
			counter = 0
			tempAsciArray = nil
		} else {
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
	for i := 0; i < 8; i++ {
		for _, v := range tempIntArrLetter {
			fmt.Print(asciiTemplates[v][i])
		}
		fmt.Println("$")
	}
}

func returnAsciiCodeInt(s string) []int {
	var tempIntArrLetter []int
	for _, v := range s {
		tempIntArrLetter = append(tempIntArrLetter, (int(v) - 32))
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
		idx := strings.Index(text, escapedN)
		if idx == -1 {
			substrings = append(substrings, text)
			break
		}

		substrings = append(substrings, text[:idx])

		if idx+len(escapedN) < len(text) && text[idx+len(escapedN)] == 'n' {
			substrings = append(substrings, newline)
			text = text[idx+len(escapedN)+1:]
		} else {
			substrings = append(substrings, escapedN)
			text = text[idx+len(escapedN):]
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
