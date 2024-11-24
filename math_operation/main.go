package main

import "fmt"

func main() {
	a := 5
	b := 10
	c := a + b

	fmt.Println(c)

	increment := c
	increment++
	fmt.Println(increment)

	decrement := c
	decrement--
	fmt.Println(decrement)

	assignIncrement := increment
	assignIncrement += increment
	fmt.Println(assignIncrement)

	assignDecrement := decrement
	assignDecrement -= decrement
	fmt.Println(assignDecrement)
}
