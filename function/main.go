package main

import "fmt"

func main() {
	sayHello()
	subtract()
	sayHelloTo("Ayom", "Izzuddin")

	mandiri := getPaymentMethod("Mandiri")
	fmt.Println(mandiri)

	bcaSyariah := getPaymentMethod("BCA Syariah")
	fmt.Println(bcaSyariah)

	paymentMethods := getListPaymentMethod()
	fmt.Println(paymentMethods)

	firstName, middleName, lastName := getFullName()
	fmt.Printf("Full Name: %s %s %s\n", firstName, middleName, lastName)

	firstName, middleName, _ = getFullName()
	fmt.Println(firstName, middleName)

	firstName1, middleName1, lastName1 := getCompleteName()
	fmt.Printf(firstName1, middleName1, lastName1)
}

func sayHello() {
	fmt.Println("Hello")
}

func subtract() {
	fmt.Println(10 - 5)
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getPaymentMethod(payment string) string {
	return "Bank " + payment
}

func getListPaymentMethod() []string {
	return []string{"Bank Mandiri", "Bank BCA", "Bank BRI"}
}

func getFullName() (string, string, string) {
	return "Muhammad", "Ayom", "Izzuddin"
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Ayom"
	lastName = "Izzuddin"
	return firstName, middleName, lastName
}
