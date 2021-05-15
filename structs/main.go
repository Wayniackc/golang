package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	// Can also just do it as contactInfo
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// The above method works, but it is dependent upon the order of the variables in the struct
	// If the order was to change at some point it would break everything
	//
	// Second method below assigns the properties to the values
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	//
	// Third method using var.field
	// var alex person
	//
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()
}
func (pointerToPerson *person) updateName(newFirstName string) {

	(*pointerToPerson).firstName = newFirstName

	// Turn address into value with *address
	// Turn value into address with &value

}

func (p person) print() {

	fmt.Printf("%+v", p)

}
