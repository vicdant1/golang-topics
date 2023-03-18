package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// dealing with errors:
	res, err := http.Get("http://google.com.br")

	if err != nil {
		log.Fatal("Erro: " + err.Error())
		// panic("error")
	}

	fmt.Println(res.Header)

}

func Soma(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("total maior que 10")
	}

	return res, nil
}
