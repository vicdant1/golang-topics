package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Using OS package and reading commands")
	fmt.Printf("%d arguments passed \n", len(os.Args))

	for index, arg := range os.Args {
		fmt.Printf("%d: %s\n", index, string(arg))
	}
}
