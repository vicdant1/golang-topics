package main

import "fmt"

func main() {
	// spread operator in go is not ...slice but slice...
	test := []int{1, 1, 23}
	fmt.Println(test)

	var slice2 []int
	slice2 = append(slice2, test...)
	fmt.Println(slice2)

	// using make function

	// make: type, len, cap
	// define cap avoids program to try to reallocate mem which is going to be important to keep program safe and fast
	testSlice := make([]int, 3, 9)

	fmt.Printf("Len: %d | Cap: %d\n", len(testSlice), cap(testSlice))

	// Map -> keyvaluepair {"Key": "Value"}
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"Adam": 23, "Sarah": 33}
	fmt.Println(myMap2["Adam"])

	age, ok := myMap2["Adam"]
	if ok {
		fmt.Printf("Adam's age is %d\n", age)
	} else {
		fmt.Printf("Name not found on map\n")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %s Age: %d\n", name, age)
	}

}
