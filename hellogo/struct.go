package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Costumer struct {
	Name   string
	Email  string
	IdCard string
}

func (costumer Costumer) PrintCostumerData() {
	fmt.Printf("Name: %s\nEmail: %s\nId Card: %s\n", costumer.Name, costumer.Email, costumer.IdCard)
}

func (internationalCostumer InternationalCostumer) PrintCostumerData() {
	// Note that I can use internationalCostumer.Name instead of internationalCostumer.Costumer.Name
	fmt.Printf("\n[***** INTERNATIONAL *****]\nName: %s\nEmail: %s\nId Card: %s\nCountry: %s\n", internationalCostumer.Name, internationalCostumer.Email, internationalCostumer.IdCard, internationalCostumer.Country)
}

type InternationalCostumer struct {
	Costumer        // Kindof composition
	Country  string `json:"country"`
	Name     string
}

func main() {
	costumer := Costumer{
		Name:   "John",
		Email:  "john@johnices.com",
		IdCard: "90123.123-11",
	}
	costumer.PrintCostumerData()

	internationalCostumer := InternationalCostumer{
		Costumer: costumer,
		Country:  "USA",
		Name:     "OVERLOADNAME",
	}

	internationalCostumer.PrintCostumerData() // intersting feature

	costumerJson, err := json.Marshal(internationalCostumer)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(costumerJson))

	jsonNewCostumer := `{"Email":"john@johnices.com","IdCard":"90123.123-11","country":"USA","Name":"John Mikalbek"}`

	newCostumer := InternationalCostumer{}
	newCostumer.PrintCostumerData()

	json.Unmarshal([]byte(jsonNewCostumer), &newCostumer)

	newCostumer.PrintCostumerData()
}
