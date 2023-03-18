package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	resultado, err := Soma(0, 1)

	if err != nil {
		log.Fatal("Não foi possível completar a operação: " + err.Error())
	}

	// Same behavior:
	// fmt.Println(fmt.Sprintf("No problem with operation, result is %d", resultado))
	fmt.Printf("No problem with operation, result is %d\n", resultado)

	noValueResult := SpecialMultiplication(71)
	fmt.Printf("Result: %d\n", noValueResult)
}

func Soma(n1 int, n2 int) (int, error) {

	soma := n1 + n2

	if soma > 10 {
		return soma, errors.New("soma maior que 10")
	}

	return n1 + n2, nil
}

func SpecialMultiplication(value int) (result int) {
	value *= 836
	result = value
	return
}
