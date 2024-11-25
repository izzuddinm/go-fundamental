package main

import "fmt"

func main() {
	blacklist := func(name string) bool {
		return name == "Izzuddin"
	}

	isBlockedUser := isBlocked("Izzuddin", blacklist)
	fmt.Println("Is user blocked:", isBlockedUser)

	isBlocked("Izzuddin", func(name string) bool {
		return name == "Izzuddin"
	})
}

type Blacklist func(string) bool

func isBlocked(name string, blacklist Blacklist) bool {
	if blacklist(name) {
		fmt.Println("You are blocked")
		return true
	} else {
		fmt.Println("Successfully registered")
		return false
	}
}
