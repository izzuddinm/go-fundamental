package main

import (
	"fmt"
)

type Customer struct {
	FirstName   string
	MiddleName  string
	LastName    string
	Email       string
	Address     string
	PhoneNumber string
	Age         int
}

func main() {
	var customer Customer
	fmt.Println(customer)

	customer.FirstName = "Muhammad"
	customer.MiddleName = "Ayom"
	customer.LastName = "Izzuddin"
	customer.Email = "izzuddin@gmail.com"
	customer.PhoneNumber = "081234567567"
	customer.Age = 20

	fmt.Println(customer)

	fmt.Println(customer.FirstName + " " + customer.MiddleName + " " + customer.LastName)

	user1 := formatFullName(customer)
	fmt.Println(user1)

	// another way
	user2 := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Ayom",
		LastName:   "Izzuddin",
	}

	fmt.Println(user2)

	user1.sayHello("Ayom")
	user2.sayHello("John")
}

func formatFullName(customer Customer) Customer {
	var cust Customer
	cust.FirstName = customer.FirstName
	cust.MiddleName = customer.MiddleName
	cust.LastName = customer.LastName

	if customer.Address == "" {
		fmt.Println("Address is null.")
	}

	return cust
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, "My name is", customer.LastName)
}
