package main

import "fmt"

type Costumer struct {
	Name   string
	Email  string
	IdCard string
}

func main() {
	costumer := Costumer{
		Name:   "John",
		Email:  "john@johnices.com",
		IdCard: "90123.123-11",
	}

	fmt.Printf("Name: %s\nEmail: %s\nId Card: %s\n", costumer.Name, costumer.Email, costumer.IdCard)
}
