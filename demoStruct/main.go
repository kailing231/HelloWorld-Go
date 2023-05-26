package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	bob := person{
		firstName: "Bob",
		lastName:  "Someone",
		contactInfo: contactInfo{
			email:   "bob@mail.com",
			zipcode: 5555,
		},
	}

	// Go implicit understands
	bob.setFirstName("bobby")
	bob.print()
}

// Set first name of person
func (p *person) setFirstName(newFirstName string) {
	p.firstName = newFirstName
}

// Print all contents of person
func (p person) print() {
	fmt.Printf("%+v", p)
}
