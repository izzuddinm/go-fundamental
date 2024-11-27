package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) MethodWithoutPointer() {
	man.Name = "Hello, Mr. " + man.Name
}

func (man *Man) MethodWithPointer() {
	man.Name = "Hello, Mr. " + man.Name
}

func main() {
	user := Man{"Muhammad Ayom Izzuddin"}
	user.MethodWithoutPointer()
	fmt.Println(user)

	user.MethodWithPointer()
	fmt.Println(user)
}
