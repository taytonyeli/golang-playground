package main

import "fmt"

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
	tony := person{
		firstName: "Tony",
		lastName:  "Tay",
		contactInfo: contactInfo{
			email:   "tt@example.com",
			zipCode: 12345,
		},
	}
	tony.updateName("Tonye")
	tony.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pp *person) updateName(newFirstName string) {
	(*pp).firstName = newFirstName
}
