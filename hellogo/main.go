package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	result, err := Sum(0, 1)

	if err != nil {
		log.Fatal("Operation problem: " + err.Error())
	}

	// Same behavior:
	// fmt.Println(fmt.Sprintf("No problem with operation, result is %d", resultado))
	fmt.Printf("No problem with operation, result is %d\n", result)

	noValueResult := SpecialMultiplication(71)
	fmt.Printf("Result: %d\n", noValueResult)

	SequencialSum(1, 2, 3, 3, 43, 42, 34, 23, 42, 34, 23, 42, 342, 3, 42)
}

func Sum(n1 int, n2 int) (int, error) {
	sum := n1 + n2
	if sum > 10 {
		return sum, errors.New("Bigger than 10")
	}

	return n1 + n2, nil
}

func SpecialMultiplication(value int) (result int) {
	value *= 836
	result = value
	return
}

func SequencialSum(x ...int) {
	fmt.Printf("%d elements", len(x))
}
