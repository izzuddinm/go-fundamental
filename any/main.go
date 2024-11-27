package main

import (
	"fmt"
)

func AnyDataTypeOrInterfaceEmpty() interface{} {
	// return 10
	// return true
	return "Oops, unknown error your application!"
}

func Logging() {
	fmt.Println("System running.")
	message := recover()
	fmt.Println("Happen error", message)
}

func main() {
	defer Logging()
	result := AnyDataTypeOrInterfaceEmpty()
	fmt.Println(result)
	panic("Error")

}
