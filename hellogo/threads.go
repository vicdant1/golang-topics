package main

import (
	"fmt"
	"time"
)

func Counter(systype string) {
	for i := 0; i < 5; i++ {
		fmt.Println(systype, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go Counter("a ")
	go Counter("b ")

	time.Sleep(time.Second * 10)
}
