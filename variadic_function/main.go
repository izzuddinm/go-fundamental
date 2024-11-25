package main

import "fmt"

func main() {
	total := sumAll(5, 4, 3, 2, 1)
	fmt.Println(total)

	total = sumAll(10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	fmt.Println(total)

	numbers := []int{10, 20, 30, 40, 50}
	total = sumAll(numbers...)
	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
