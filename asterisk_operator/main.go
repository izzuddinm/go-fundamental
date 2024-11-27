package main

import "fmt"

type Address struct {
	City    string
	Provice string
	Country string
}

func main() {
	address3 := Address{"Banyuwangi", "Jawa Timur", "Indonesia"}
	address4 := &address3

	address3.City = "Tuban"

	fmt.Println(address3)
	fmt.Println(address4)

	*address4 = Address{}
	fmt.Println(address3)
	fmt.Println(address4)
}
