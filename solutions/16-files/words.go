package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Create an empty map. It will be used to map words to counts.
	wordcount := make(map[string]int)

	if len(os.Args) < 2 {
		fmt.Printf("usage: %s filename", os.Args[0])
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	// We'll use a regular expression to remove punctuation. This
	// regular expression matches any group of characters which are
	// not (lower case) letters or spaces. So any punctuation will
	// not match.
	reg, err := regexp.Compile("[^a-z ]+")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// Take the line returned by the scanner and make it lowercase
		line := strings.ToLower(scanner.Text())
		// Use the regular expression to convert all sequences of non-
		// letters (as well as spaces) to nothing, i.e., remove them
		line = reg.ReplaceAllString(line, "")
		// Split the line into words
		words := strings.Split(line, " ")

		for _, word := range words {
			wordcount[word]++
		}
	}
	for key, val := range wordcount {
		fmt.Println(key, val)
	}
}
