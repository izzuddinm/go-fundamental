package main

import "fmt"

func main() {
	var names = [3]string{"Muhammad", "Ayom", "Izzuddin"}

	for i := 0; i < len(names); i++ {
		index := fmt.Sprintf("Index on %d with value %s", i, names[i])
		fmt.Println(index)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
