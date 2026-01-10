package main

import (
	"fmt"
	"project/magazine"
)

func main() {
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
	subscriber2 := magazine.Subscriber{Name: "Jack", Rate: 120.0, Active: true}
	fmt.Println(subscriber2)
	// you can init struct with only couple of fields for example:
	subscriber3 := magazine.Subscriber{Rate: 1000}
	fmt.Println(subscriber3)
	employee := magazine.Employee{
		Name:   "Jackson",
		Salary: 499.99,
		HomeAddress: magazine.Address{
			Street:   "13th Avenue",
			City:     "NY",
			State:    "NY",
			Postcode: 12345,
		},
		ContactInfo: magazine.ContactInfo{
			PhoneNumber: "19384323",
			Email:       "employee@gmail.com",
		}, // but directly you need to call it like this.
	}
	fmt.Printf("%s, salary is: %2.f\n", employee.Name, employee.Salary)
	fmt.Println(employee)
	fmt.Println(employee.HomeAddress.City) // it's ambiguous
	fmt.Println(employee.PhoneNumber)      // with anonymous style I can call them directly
}
