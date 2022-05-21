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
	p.updateName("Joao")
	p.updateEmail("joao@gmail.com")

	p.print()

}

func (p *person) updateName(fName string) {
	(*p).FirstName = fName
}

func (p *person) updateEmail(e string) {
	(*p).Email = e
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
