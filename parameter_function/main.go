package main

import "fmt"

func main() {
	sayHelloWithFilter("Dog", spamFilter)
	sayHelloWithFilter("Muhammad Ayom Izzuddin", spamFilter)
	sayHelloWithFilterParameter("Muhammad Ayom Izzuddin", spamFilter)
}

type Filter func(name string) string

func sayHelloWithFilter(name string, filter func(name string) string) {
	filteredName := filter(name)
	fmt.Println(filteredName)
}

func sayHelloWithFilterParameter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println(filteredName)
}

func spamFilter(name string) string {
	if name == "Dog" {
		return "..."
	} else if name == "Anjing" {
		return "change your name!"
	} else {
		return name
	}
}
