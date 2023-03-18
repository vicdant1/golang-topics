package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	res, err := Soma(1, 1)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res)
}

func Soma(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("total maior que 10")
	}

	return res, nil
}
