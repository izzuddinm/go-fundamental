package main

import "fmt"

func main() {
	type Name string

	var fullName Name = "Muhammad Ayom Izzuddin"
	fmt.Println(fullName)

	var firstIndexFullNameString = Name(fullName[0])
	fmt.Println(firstIndexFullNameString)
}
