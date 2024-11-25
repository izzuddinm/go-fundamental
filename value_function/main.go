package main

import "fmt"

func main() {
	sampleOne := getGoodBye("Muhammad Ayom Izzuddin")
	fmt.Println(sampleOne)

	sampleTwo := getGoodBye

	value := sampleTwo("Muhammad Ayom Izzuddin")
	fmt.Println(value)
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
