package main

import (
	"fmt"
	"time"
)

func Counter(systype string) {
	for i := 0; i < 10; i++ {
		fmt.Println(systype, i)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Starting execution")

	go Counter("goroutine")
	for i := 0; i < 5; i++ {
		fmt.Println(i, "<=>")
		time.Sleep(time.Second * 1 / 2)
	}

	time.Sleep(time.Second * 10)
}
