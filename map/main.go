package main

import "fmt"

func main() {
	var personOne map[string]string = map[string]string{}
	personOne["name"] = "Muhammad Ayom Izzuddin"
	personOne["country"] = "Indonesia"
	fmt.Println(personOne)
	fmt.Println(personOne["name"])
	fmt.Println(personOne["country"])

	personTwo := map[string]string{
		"name":    "Cristiano Ronaldo",
		"country": "Portugal",
	}
	fmt.Println(personTwo)
	fmt.Println(personOne["name"])
	fmt.Println(personOne["country"])

	// function map
	lenPersonOne := len(personOne)
	fmt.Println(lenPersonOne)

	book := make(map[string]string)
	book["title"] = "Golang Fundamental"
	book["Author"] = "Muhammad Ayom Izzuddin"
	book["False"] = "False"

	fmt.Println(book)
	delete(book, "False")
	fmt.Println(book)
}
