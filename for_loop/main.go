package main

import "fmt"

func main() {
	counter := 0
	var names = []string{"Muhammad", "Ayom", "Izzuddin"}

	for counter < 10 {
		fmt.Println("Index on ", counter)
		counter++
	}

	for i := 0; i < len(names); i++ {
		index := fmt.Sprintf("Index on %d with value %s", i, names[i])
		fmt.Println(index)
	}

	for index, name := range names {
		fmt.Println("Index = ", index, "Name = ", name)
	}

	for _, name := range names {
		fmt.Println("Name = ", name)
	}
}
