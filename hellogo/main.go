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

	// DivideScreen()

	// UnderstandingFors()
}

func DivideScreen() {
	fmt.Println("*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=")
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
	fmt.Printf("%d elementsin sequence\n", len(x))
	var sum int
	for _, element := range x {
		fmt.Printf("%d ", element)
		sum += element
	}
	fmt.Println()

	fmt.Printf("The sum of the %d elements is %d", len(x), sum)
}

func UnderstandingFors() {
	// Usual for
	for i := 0; i <= 10; i++ {
		fmt.Printf(" %d ", i)
	}
	fmt.Println()

	DivideScreen()

	// Go while
	lowerThan10 := true
	counter := 0
	for lowerThan10 {
		fmt.Printf("Current number: %d\n", counter)
		if counter == 10 {
			lowerThan10 = false
		}
		counter++
	}

	DivideScreen()

	// Go foreach (reads key and value of a array)
	intArray := [...]int{0, 1, 2, 2, 2, 3, 4, 1}
	for key, value := range intArray {
		fmt.Printf("Key: %d, Value: %d\n", key, value)
	}

}
