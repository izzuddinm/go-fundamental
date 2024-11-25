package main

import "fmt"

func main() {
	var a, b int = 5, 8
	var c, d = "Java", "Golang"

	total := a + b
	fmt.Println(total)

	if total == a+b {
		fmt.Println("isValid")
	} else {
		fmt.Println("isNotValid")
	}

	if c == d {
		fmt.Println("isValid")
	} else if c == "Java" {
		fmt.Println("isValid")
	} else {
		fmt.Println("isNotValid")
	}

	time := 23
	if time < 10 {
		fmt.Println("Hai Good morning.")
	} else if time < 20 {
		fmt.Println("Hai Good day.")
	} else {
		fmt.Println("Hai Good evening.")
	}

	name := "Muhammad Ayom Izzuddin"
	lenNameOne := len(name)
	if lenNameOne > 10 {
		fmt.Println("Name Too Long!")
	} else {
		fmt.Println("Format Name Is Correct")
	}

	// can use if short statement
	if lenNameTwo := len(name); lenNameTwo > 10 {
		fmt.Println("Name Too Long!")
	} else {
		fmt.Println("Format Name Is Correct")
	}
}
