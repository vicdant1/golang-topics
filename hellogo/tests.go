package main

import "fmt"

func InfinitLoop() {
	counter := 0
	for {
		fmt.Printf(" %d ", counter)
		counter++
	}
}

func NumberPrinting(ch chan<- bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d ", i)
	}

	ch <- true
}

func main() {
	go InfinitLoop()

	ch := make(chan bool)

	go NumberPrinting(ch)

	<-ch
}
