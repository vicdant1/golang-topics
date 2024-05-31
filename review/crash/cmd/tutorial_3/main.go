package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe()

	printMeParam("Hello, world as well!")

	numerator := 11
	denominator := 2

	var result, remainder, err = intDivision(numerator, denominator)

	// if-else statement
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %d\n", result)
	} else {
		fmt.Printf("The result of the integer division is %d with remainder %d\n", result, remainder)
	}

	// switch is very flexible in go
	// this structure works like a if/else if/else block

	// btw switch statement has implicit break inside on this format
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %d\n", result)
	default:
		fmt.Printf("The result of the integer division is %d with remainder %d\n", result, remainder)
	}

	// conditional switch case
	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}

	variadicFunction(1, 2, 3, 123, 1, 234, 12, 341, 2, 3465, 34, 53, 1243, 3)
	test := namedReturnFunction()
	fmt.Println(test)
}

// by default golang exported functions should use PascalCase and private functions should use camelCase since we have no public or private declarations in GO
func printMe() {
	fmt.Println("Hello, world!")
}

func printMeParam(printValue string) {

	fmt.Println(printValue)

}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil
}

func variadicFunction(numbers ...int) {
	var sum int
	for _, x := range numbers {
		sum += x
	}

	fmt.Printf("Total amount of params: %d\nResult of this sum: %d\n", len(numbers), sum)
}

func namedReturnFunction() (x int) {
	x = 2
	return
}
