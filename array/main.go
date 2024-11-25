package main

import "fmt"

func main() {
	var programmingLanguages [3]string

	programmingLanguages[0] = "Java"
	programmingLanguages[1] = "Golang"
	programmingLanguages[2] = "Ruby On Rails"

	fmt.Println(programmingLanguages[0])
	fmt.Println(programmingLanguages[1])
	fmt.Println(programmingLanguages[2])

	var numbers = [3]int{3, 2, 1}
	fmt.Println(numbers)

	var dataTypeGolang = [3]string{
		"int",
		"string",
		"float",
	}
	fmt.Println(dataTypeGolang)

	var names = [...]string{
		"Muhammad",
		"Ayom",
		"Izzuddin",
	}
	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(len(names))
}
