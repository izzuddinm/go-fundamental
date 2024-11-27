package main

import "fmt"

type address string

type Address struct {
	City    address
	Provice address
	Country address
}

func main() {
	/*
		address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
		fmt.Println(address1)

		address2 := address1
		fmt.Println(address2)

		// address1 copy to address2

		address1.City = "Surabaya"
		address1.Provice = "Jawa Timur"
		fmt.Println(address1)
		fmt.Println(address2)
	*/

	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	address2 := &address1
	fmt.Println(address2) // pointer use &

	address1.City = "Surabaya"
	address1.Provice = "Jawa Timur"
	fmt.Println(address1)
	fmt.Println(address2)
}
