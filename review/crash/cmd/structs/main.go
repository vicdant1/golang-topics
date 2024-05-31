package main

import (
	"fmt"
	"strings"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) calcMilesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func (e gasEngine) calcMilesLeft() uint8 {
	return e.gallons * e.mpg
}

type engine interface {
	calcMilesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles < e.calcMilesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	myEngine := gasEngine{
		25,
		15,
	}

	myEngine.mpg = 20

	fmt.Println(myEngine)
	fmt.Println(myEngine.mpg, myEngine.gallons)

	// how do GO deals with anonymous structs? (Like dynamic)
	myEngine2 := struct {
		mpg     uint8
		gallons uint8
	}{25, 15}

	fmt.Println(myEngine2)
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	fmt.Println(strings.Repeat("=", 30))

	fmt.Printf("Total miles left in tank %d\n", myEngine.calcMilesLeft())

	fmt.Println(strings.Repeat("=", 30))

	canMakeIt(myEngine, 22)
	canMakeIt(myEngine, 88)
}
