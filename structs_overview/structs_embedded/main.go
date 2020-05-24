package main

import "fmt"

type contact struct {
	email, address string
	phone          int
}

type employee struct {
	name        string
	salary      int
	contactInfo contact
}

func main() {
	john := employee{
		name:   "John Keller",
		salary: 4000,
		contactInfo: contact{
			email:   "jkeller@company.com",
			address: "street 20, London",
			phone:   8014654833,
		},
	}

	fmt.Printf("%+v\n", john)
	fmt.Printf("Employees's email: %s\n", john.contactInfo.email)

	john.contactInfo.email = "new_email@company.com"
	fmt.Printf("Employee's new email: %s\n", john.contactInfo.email)

	myContact := contact{email: "test@gmail.com", phone: 5445699775, address: "Texas, United States"}

	fmt.Printf("my contact info: %+v\n", myContact)

}
