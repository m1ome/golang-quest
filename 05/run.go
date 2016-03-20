package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Quest: Count word occurencies in a text file")

	var word string
	fmt.Println("Print a word to search:")
	fmt.Scanln(&word)

	fmt.Printf("Scanning file for word \"%s\"\n", word)

	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println("Error: can't open 'input.txt' in quest folder")
	}

	defer file.Close()

	// Lowercasing input word
	wordSearch := strings.ToLower(word)

	occurencies := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, w := range strings.Split(scanner.Text(), " ") {
			if matcher := strings.ToLower(w); matcher == wordSearch {
				occurencies = occurencies + 1
			}
		}
	}

	if occurencies > 0 {
		fmt.Printf("Total occurencies of \"%s\" is %d\n", word, occurencies)
	} else {
		fmt.Printf("We can't find any occurencies of \"%s\"\n", word)
	}
}
