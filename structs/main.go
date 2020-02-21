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

	// khalil := person{"Khalil", "Mack"}

	//protects against re-organization of the person stuct's fields
	// khalil := person{firstName: "Khalil", lastName: "Mack"}

	//here a zero value is assigned to the struct
	//meaning all its fields are assigned to their defaults: ""
	var khalil person
	fmt.Println(khalil)

	khalil.firstName = "Khalil"
	khalil.lastName = "Mack"
	//%+v will print out the struct with its named args
	fmt.Printf("%+v", khalil)

	jugg := person{
		firstName: "The",
		lastName:  "Juggurnaut",
		contactInfo: contactInfo{
			email:   "the_juggarnaut@b.com",
			zipCode: 12345,
		},
	}
	jugg.print()

	bal := person{
		firstName: "I'm a ",
		lastName:  "Baloon",
		contactInfo: contactInfo{
			email:   "baloon@pop.com",
			zipCode: 12345,
		},
	}

	bal.updateName("Big ol' ")
	bal.print()
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	fmt.Printf(">>>> \n %v", pointerToPerson)
	(*pointerToPerson).firstName = newFirstName
}
