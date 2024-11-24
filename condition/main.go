package main

import "fmt"

func main() {
	var a, b int = 5, 8

	total := a + b
	fmt.Println(total)

	if total == a+b {
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
}
