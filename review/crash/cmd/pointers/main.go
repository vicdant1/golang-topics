package main

import (
	"fmt"
	"math"
)

func main() {
	// allocates a mem add to a "zero" value of passed type
	var p *int32 = new(int32)
	*p = 10
	var i int32
	fmt.Printf("i: %d, p: %d\n", i, *p)
	fmt.Printf("Addresses information =======\ni:%v, p:%v\n", &i, p)
	p = &i
	fmt.Println("p pointing to i")
	fmt.Printf("i: %d, p: %d\n", i, *p)
	fmt.Printf("Addresses information =======\ni:%v, p:%v\n", &i, p)

	powSlice := []int{1, 2, 3, 4}
	powMembers(powSlice)

	fmt.Printf("Actual slice: %v\n", powSlice)

}

func powMembers(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = int(math.Pow(float64(slice[i]), float64(2)))
	}

	fmt.Printf("Current slice: %v\n", slice)
}
