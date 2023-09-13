// this is our group project from chatgpt
package main

import (
	"fmt"
	"os"
)

// Define a map to store ASCII representations of characters
var bannerMap = make(map[rune]string)

// Load banner templates from files
func loadBannerTemplates() {
	banners := map[string]string{
		"shadow":     "https://learn.01founders.co/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt",
		"standard":   "https://learn.01founders.co/git/root/public/src/branch/master/subjects/ascii-art/standard.txt",
		"thinkertoy": "https://learn.01founders.co/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt",
	}

	for name, url := range banners {
		resp, err := fetchBannerTemplate(url)
		if err != nil {
			fmt.Printf("Error loading banner template %s: %v\n", name, err)
			os.Exit(1)
		}
		bannerMap[rune(name[0])] = resp
	}
}

// Fetch banner template from a URL
func fetchBannerTemplate(url string) (string, error) {
	resp, err := os.ReadFile(url)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// Convert a character to its ASCII representation
func charToASCII(c rune) string {
	if ascii, ok := bannerMap[c]; ok {
		return ascii
	}
	return ""
}

// Convert a string to its ASCII representation
func stringToASCII(input string) string {
	var result string
	for _, c := range input {
		if ascii := charToASCII(c); ascii != "" {
			result += ascii + "\n"
		} else {
			result += "\n" // Use a newline for unsupported characters
		}
	}
	return result
}

func main() {
	// Load banner templates
	loadBannerTemplates()

	// Accept input string from command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run asciiart.go <input_string>")
		os.Exit(1)
	}
	input := os.Args[1]

	// Convert and print the input string as ASCII art
	asciiArt := stringToASCII(input)
	fmt.Println(asciiArt)
}
