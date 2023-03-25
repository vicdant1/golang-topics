package main

import "fmt"

func main() {

	fmt.Println("Printing first")
	defer fmt.Println("Printing second")
	fmt.Println("Printing third")
}
