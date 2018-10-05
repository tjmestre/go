package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "jim",
		lastName:  "edg",
		contactInfo: contactInfo{
			email:   "fdgg@fsdfsdg.com",
			zipCode: 2333,
		},
	}

	jim.updateName("2222")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
