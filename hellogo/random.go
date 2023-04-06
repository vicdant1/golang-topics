package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixMicro())

	for {
		// generate random numbers
		randNumber := rand.Intn(1001)

		fmt.Println(randNumber)

		time.Sleep(time.Second / 2)
	}
}
