package main

import "fmt"

func main() {
	person := Person{
		Name:  "Muhammad Ayom Izzuddin",
		Email: "izzuddin@gmail.com",
	}

	fmt.Println(person)

	SayHello(person)
	SetStruct(person)

	animal := Animal{
		Name: "Cat",
	}

	SayHello(animal)
	SetStruct(animal) // Call SetStruct with animal
}

type Person struct {
	Name  string
	Email string
}

type Animal struct {
	Name string
}

// Define the HasName interface
type HasName interface {
	GetName() string
	GetStruct() Person
}

// Implement GetName for Person
func (person Person) GetName() string {
	return person.Name
}

// Implement GetName for Animal
func (animal Animal) GetName() string {
	return animal.Name
}

// Implement GetStruct for Person
func (person Person) GetStruct() Person {
	return person // Return the instance itself
}

// Implement GetStruct for Animal
func (animal Animal) GetStruct() Person {
	return Person{Name: animal.Name} // Create a Person from Animal
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func SetStruct(value HasName) {
	fmt.Println("Struct:", value.GetStruct())
}
