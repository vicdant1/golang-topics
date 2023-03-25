package main

import (
	"bufio"
	"fmt"
	"os"
)

// this is an algorithm to catch duplicated words

func main() {
	fmt.Println("Dup Algorithm")

	// string: int
	// "brasilia": 2
	words := make(map[string]int)

	file, err := os.Open("../static/duplicated.txt")
	if err != nil {
		fmt.Println(err.Error())
		panic("An error occured")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words[scanner.Text()]++
	}

	for city, occurrences := range words {
		fmt.Printf("%s: %d\n", city, occurrences)
	}
}
