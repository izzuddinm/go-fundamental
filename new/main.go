package main

import "fmt"

type Address struct {
	City    string
	Provice string
	Country string
}

func main() {
	var address5 *Address = new(Address)
	var address6 *Address = address5

	address6.City = "Manado"

	fmt.Println(address5)
	fmt.Println(address6)

}
