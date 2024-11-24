package main

import "fmt"

func main() {
	var firstName string = "Muhammad"
	fmt.Println(firstName)

	var middleName = "Ayom"
	fmt.Println(middleName)

	lastName := "Izzuddin"
	fmt.Println(lastName)

	var lenFirstName int = len(firstName)
	fmt.Println(lenFirstName)

	var lenMiddleName = len(middleName)
	fmt.Println(lenMiddleName)

	lenLastName := len(lastName)
	fmt.Println(lenLastName)

	var (
		fullName string = "Muhammad Ayom Izzuddin"
		email           = "izzuddinm@gmail.com"
	)

	fmt.Println(fullName)
	fmt.Println(email)
}
