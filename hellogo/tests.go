package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

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

func Worker(identifier string, ch <-chan int) {
	for v := range ch {
		fmt.Printf("[WORKER %s] Number %d\n", identifier, v)
		time.Sleep(time.Second)
	}
}

func testGoRoutine() {
	i := 0
	for i < 5 {
		fmt.Println("Test")
		time.Sleep(time.Second)
		i++
	}

}

func main() {
	// go InfinitLoop()
	// ch := make(chan bool)
	// go NumberPrinting(ch)
	// <-ch

	// ch := make(chan int)

	// go Worker("1", ch)
	// go Worker("2", ch)

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("Escrevendo %d\n", i)
	// 	ch <- i
	// }

	// close(ch)

	var teste string
	scanner := bufio.NewScanner(os.Stdin)

	// scanner.Scan()
	teste = scanner.Text()

	fmt.Println(teste)

	fmt.Println("=*=*=*=*=*=*=*=*| FOR WITH STEP |*=*=*=*=*=*=*=*=")

	for i := 0; i <= 40; i += 5 {
		fmt.Printf("Current index: %d\n", i)
	}

	// consts

	const pi = 3.14
	const tax = 0.02

	fmt.Printf("%v %v\n", pi, tax)
	go testGoRoutine()
	panic("test")
}
