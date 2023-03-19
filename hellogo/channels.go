package main

import (
	"fmt"
	"time"
)

func Write(identifier string) {
	for i := 0; i < 10; i++ {
		fmt.Printf(" %s - %d \n", identifier, i)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func main() {
	// Channels works in order to make the PIPE scheme, it allows a thread to communicate with others

	go Write("a")
	go Write("b")

	time.Sleep(time.Second * 20)
}
