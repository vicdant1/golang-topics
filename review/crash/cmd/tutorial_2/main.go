package main

import "fmt"

func main() {
	var intNum int16
	fmt.Println(intNum)

	var floatNum float64
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2

	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var myString string = "Hello, World"

	fmt.Println(myString)

	var myOtherString = `abluble
	
	
	abluble abluble`

	fmt.Println(myOtherString)

	fmt.Println(len("testaaaaaaaaaaa"))

	var testBool bool

	fmt.Println(testBool)

	// var myVar = ""
	// myVar := ""
	// var myVar string = ""
	// var myVar string
	// test1, test2, test3 := 0 // << it works fine

	const myConst = ""
	const myConst2 string = ""
	// const test := 1 // It wont work

}
