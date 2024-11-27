package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func changeAddressWithoutPointer(address Address) {
	address.Country = "Indonesia"
}

func changeAddressWithPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Surabaya", "Jawa Timur", ""}

	changeAddressWithoutPointer(address)
	fmt.Println(address)

	changeAddressWithPointer(&address)
	fmt.Println(address)
}
