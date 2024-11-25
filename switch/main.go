package main

import "fmt"

func main() {
	day := 9

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a Weekday")
	}

	name := "Muhammad Ayom Izzuddin"
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name Too Long")
	case false:
		fmt.Println("Format Name Is Correct")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name Too Long")
	case length < 1:
		fmt.Println("Name Cannot Be Null")
	default:
		fmt.Println("Format Name Is Correct")
	}
}
