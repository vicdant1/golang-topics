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

type Vehicle interface {
	Buzz()
}

type Motocicle struct {
	name  string
	brand string
}

func (moto Motocicle) Buzz() {
	fmt.Printf("ZUUUUUUUUUUUUUUUUUUUUUUUUM [%s, %s]\n", moto.name, moto.brand)
}

type Person struct {
	name string
	age  int
}

type FisicalPerson struct {
	Person
	idCard  string
	surname string
	vehicle Vehicle
}

func (fp FisicalPerson) PrintInformation() {
	fmt.Println(fp.name, fp.surname, fp.age, fp.idCard)
	fp.vehicle.Buzz()
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

	moto := Motocicle{
		name:  "CG FAN 125",
		brand: "Honda",
	}

	fisicalPerson := FisicalPerson{
		Person: Person{
			name: "Joao",
			age:  21,
		},
		idCard:  "0340342",
		surname: "Dantas",
		vehicle: moto,
	}

	fisicalPerson.PrintInformation()

	juridicPerson := JuridicPerson{
		Person: Person{
			name: "Joao LTDA",
			age:  2,
		},
		cnpj:         "92.132.1231/4214-22",
		socialReason: "JOAOZINHO LTDA",
	}

	// ShowData(fisicalPerson)
	// ShowData(juridicPerson)
	// InterfaceInformation(fisicalPerson)
	AssertType(fisicalPerson)
	AssertType(juridicPerson)
}

func AssertType(p IPerson) {
	// one by one, cannot validate many types in a single call
	if item, ok := p.(FisicalPerson); ok {
		fmt.Printf("%s is a Fisical Person\n", item.name)
	}

	if item, ok := p.(JuridicPerson); ok {
		fmt.Printf("%s is a Juridic Person and its CNPJ is %s\n", item.name, item.cnpj)
	}

	// many types validation
	switch p.(type) {
	case FisicalPerson:
		fmt.Println("Fisical Person")
	case JuridicPerson:
		fmt.Println("Juridic Person")
	default:
		fmt.Println("Unreconized Type")
	}

	fmt.Println()
}
