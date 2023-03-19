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

func ProcessNumbers(numberChannel chan<- int) {
	fmt.Println("Entrou Process")
	for i := 0; i < 10; i++ {
		fmt.Printf("Escrevendo %d\n", i)
		numberChannel <- i
	}
	fmt.Println("Saiu Process")

	close(numberChannel)
}

func main() {
	// Channels works in order to make the PIPE scheme, it allows a thread to communicate with others

	// cn1 := make(chan bool)
	// cn2 := make(chan bool)

	// go Write("a", cn1)
	// go Write("b", cn2)

	// Making main (wich is a go routine as well) be blocked til cn1 & cn2 have a state change
	// <-cn1
	// <-cn2
	fmt.Println("entrando main")

	numberChannel := make(chan int)

	go ProcessNumbers(numberChannel)

	for v := range numberChannel {
		fmt.Printf("Lendo %d\n", v)
	}
	fmt.Println("saindo main")
}
