package main

import "fmt"

func main() {
	var a string = "Golang"
	var b string = "Golang"
	var c string = "Java"
	var d int = 10
	var f int = 15

	var isSame bool = a == b
	fmt.Println(isSame)

	var isDifferent bool = a == c
	fmt.Println(isDifferent)

	var isValid bool = a != c
	fmt.Println(isValid)

	var isMoreThan bool = f > d
	fmt.Println(isMoreThan)
}
