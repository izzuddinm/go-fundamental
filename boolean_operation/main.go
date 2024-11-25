package main

import "fmt"

func main() {
	var finalGrade int = 74
	var absenceScore int = 80

	if finalGrade >= 75 && absenceScore >= 80 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Not Passed")
	}
}
