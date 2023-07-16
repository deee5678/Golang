package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct { // this is a custom type named person of data structure struct with the following fields
	firstName   string //No commas here
	lastName    string
	contactInfo //contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Pai",
		//contact: contactInfo{}
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}
	/*jimPointer := &jim
	jimPointer.updateName("jimmy") */
	jim.updateName("jimmy") // the same above job will be done in this single line
	jim.print()             //fmt.Printf("%+v", jim)
}

func (pointerToPerson *person) updateName(nweFirstName string) {
	(*pointerToPerson).firstName = nweFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
} */

/* func main() {
	// Deepthi :=  person{"Deepthi","Poojary"}   // declaring initialising and assigning in one step here. hdere fname & lname should be in order. you cant swap the postion of fname & lname
	Deepthi := person{firstName: "Deepthi", lastName: "Poojary"} // here u can swap the positions
	fmt.Println(Deepthi)
} */
