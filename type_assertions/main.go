package main

import "fmt"

func Random() any {
	return "random"
}

func main() {
	value := Random()

	if convertDataToString, ok := value.(string); ok {
		fmt.Println("String value:", convertDataToString)
	} else {
		fmt.Println("Error: Value is not a string")
	}

	if convertDataToInt, ok := value.(int); ok {
		fmt.Println("Integer value:", convertDataToInt)
	} else {
		fmt.Println("Error: Value is not an integer")
	}

	switch result := value.(type) {
	case string:
		fmt.Println("Data stype of result is (string) :", result)
	case int:
		fmt.Println("Data stype of result is (int) :", result)
	case bool:
		fmt.Println("Data stype of result is (bool) :", result)
	default:
		fmt.Println("Unknown data type of result! :", result)
	}

}
