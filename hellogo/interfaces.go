package main

import "fmt"

type IPerson interface {
	PrintInformation()
}

func InterfaceInformation(person IPerson) {
	// fmt.Println(person.(type)) => it will never work, .type can only be used in switch
	fmt.Println(person)
}

func ShowData(person IPerson) {

	fmt.Printf("============= TYPE =============\n%T\n", person) // the closest we are going to have in order to verify types

	switch person.(type) {
	case JuridicPerson:
		fmt.Println("Print Juridic Person: ")
	case FisicalPerson:
		fmt.Println("Print Fisic Person: ")
	default:
		fmt.Println("Unreconized type")
		return
	}

	person.PrintInformation()
}

type Person struct {
	name string
	age  int
}

type FisicalPerson struct {
	Person
	idCard  string
	surname string
}

func (fp FisicalPerson) PrintInformation() {
	fmt.Println(fp.name, fp.surname, fp.age, fp.idCard)
}

type JuridicPerson struct {
	Person
	cnpj         string
	socialReason string
}

func (jp JuridicPerson) PrintInformation() {
	fmt.Println(jp.name, jp.socialReason, jp.age, jp.cnpj)
}

func main() {
	fisicalPerson := FisicalPerson{
		Person: Person{
			name: "Joao",
			age:  21,
		},
		idCard:  "0340342",
		surname: "Dantas",
	}

	juridicPerson := JuridicPerson{
		Person: Person{
			name: "Joao LTDA",
			age:  2,
		},
		cnpj:         "92.132.1231/4214-22",
		socialReason: "JOAOZINHO LTDA",
	}

	ShowData(fisicalPerson)
	ShowData(juridicPerson)
	InterfaceInformation(fisicalPerson)
}
