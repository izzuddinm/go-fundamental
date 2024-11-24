package main

import (
	"fmt"
)

func main() {

	// Declaring (Creating) Variables
	fmt.Println("Hello, I was learn golang")
	sentence := getSentence()
	fmt.Println(sentence)

	var numStaticType int
	numStaticType = 10
	fmt.Println(numStaticType)

	numDynamicType := 20
	fmt.Println(numDynamicType)

	itemsCount := 5
	price := 100.50
	formatedString := fmt.Sprintf("itemsCount %d - price %.2f", itemsCount, price)
	fmt.Println(formatedString)

	isValid := true
	fmt.Println(isValid)
	var isNotValid bool = false
	println(isNotValid)

	// Go Multiple Variable Declaration
	var a, b, c, d int = 1, 2, 3, 4
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var e, f = 5, "Hello World"
	fmt.Println(e, f)

	g, h := 5, "Hello World"
	fmt.Println(g, h)

	// Go Variable Declaration in a Block
	var (
		i int
		j int     = 1
		k string  = "Hello World"
		l float64 = 120.34
	)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	// Go Variable Naming Rules
	// Camel Case
	var myVariableName string = "Muhammad Ayom Izzuddin"
	fmt.Println(myVariableName)

	// Pascal Case
	var MyVariableName string = "Muhammad Ayom Izzuddin"
	fmt.Println(MyVariableName)

	// Snake Case
	var my_variable_name string = "Muhammad Ayom Izzuddin"
	fmt.Println(my_variable_name)

	// Go Contants
	const MIN_RETRY int = 3
	const MAX_RETRY int = 9
	fmt.Println(MIN_RETRY + MAX_RETRY)

	var nums = 15.5
	var txt = "Hello World!"

	// General Formatting Verbs
	fmt.Printf("%v\n", nums)
	fmt.Printf("%#v\n", nums)
	fmt.Printf("%v%%\n", nums)
	fmt.Printf("%T\n", nums)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

}
