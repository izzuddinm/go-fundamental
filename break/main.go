package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Index On ", i)
		if i == 5 {
			break
		}
	}
}
