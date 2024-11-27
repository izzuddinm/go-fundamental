package main

import "fmt"

func NewMap(name string, email string, age int8) map[string]any {
	if name == "" {
		return nil
	} else {
		return map[string]any{
			"Name":  name,
			"Email": email,
			"Age":   age,
		}
	}
}

func main() {
	newMap1 := NewMap("Muhammad Ayom Izzuddin", "izzuddinm@gmail.com", 20)
	fmt.Println(newMap1)

	newMap2 := NewMap("", "example@gmail.com", 20)
	fmt.Println(newMap2)

	if newMap1 == nil || newMap2 == nil {
		fmt.Println("newMap1 or newMap2 is nil")
	}
}
