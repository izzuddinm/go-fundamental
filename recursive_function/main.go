package main

import "fmt"

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}

func factorialLoop(count int) int {
	result := 1
	for i := count; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(count int) int {
	if count == 0 || count == 1 {
		return 1
	} else {
		return count * factorialRecursive(count-1)
	}
}
