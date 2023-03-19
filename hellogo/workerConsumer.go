package main

import "fmt"

func Worker(identifier string, ch chan<- int) {
	fmt.Printf("Entering worker %s\n", identifier)
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("Writing %d\n", i)
	}

	fmt.Printf("Exiting worker %s\n", identifier)

	close(ch)
}

func main() {
	ch := make(chan int)
	go Worker("1", ch)

	for v := range ch {
		fmt.Printf("Reading %d\n", v)
	}
}
