package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading many lines from a .csv file")

	file, err := os.Open("../static/municipios.csv")
	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	readFile(file)
}

func readFile(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		// time.Sleep(time.Millisecond * 100)
	}

}
