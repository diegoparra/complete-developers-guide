package main

import "fmt"

type contactInfo struct {
	Email   string
	ZipCode int
}

type person struct {
	FirstName string
	LastName  string
	contactInfo
}

func main() {

	p := person{
		FirstName: "Diego",
		LastName:  "Parra",
		contactInfo: contactInfo{
			Email:   "diego@gmail",
			ZipCode: 9771220,
		},
	}
	p.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
