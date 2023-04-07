package main

import "fmt"

func main() {
	// Slices x Arrays
	/*
		- Slices: slices are "unlimited size" arrays. It's like a data structure which supports values pointers
		- Arrays: usual arrays
	*/

	// it deals automatically with memory allocation
	intSlice := []int{1, 1, 3, 4, 5, 2}

	printData(intSlice)

	intArray := [3]int{1, 1, 3}

	printData(intArray[:])

	intArray[2] = 60

	printData(intArray[:])

	intSlice = append(intSlice, 100001)
	printData(intSlice)

	// var newIntSlice = []int{0, 1, 2}
	// var newIntArray [3]int
	// var newIntSliceDec []int
	// newIntSliceTest := []int{23, 3, 2, 1}
	// newIntArrayTest := [3]int{1, 2, 3}

	fmt.Println("=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*")
	// make([]type, length(how many positions placed), capacity (how many positions I have))
	testVar := make([]int, 10, 10)
	printData(testVar)

	testVar = append(testVar, 100)
	testVar = append(testVar, 100)
	testVar = append(testVar, 100)
	testVar = append(testVar, 100)
	printData(testVar)

}

func printData(data []int) {
	for _, v := range data {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
