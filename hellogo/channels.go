package main

import (
	"fmt"
	"time"
)

func Write(identifier string, channel chan<- bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf(" %s - %d \n", identifier, i)
		time.Sleep(time.Second)
	}
	fmt.Println()
	channel <- true
}

func main() {
	// Channels works in order to make the PIPE scheme, it allows a thread to communicate with others

	cn1 := make(chan bool)
	cn2 := make(chan bool)

	go Write("a", cn1)
	go Write("b", cn2)

	// Making main (wich is a go routine as well) be blocked til cn1 & cn2 have a state change
	<-cn1
	<-cn2

}
