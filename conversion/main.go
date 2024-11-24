package main

import (
	"fmt"
)

func main() {
	var valInt32 int32 = 32738
	var valInt64 int64 = int64(valInt32)
	var valInt16 int16 = int16(valInt32)
	var valInt8 int8 = int8(valInt32)

	fmt.Println(valInt32)
	fmt.Println(valInt64)
	fmt.Println(valInt16)
	fmt.Println(valInt8)

	name := "Muhammad Ayom Izzuddin"
	fmt.Println(name)

	firstIndex := name[0]
	firstIndexString := string(firstIndex)
	fmt.Println(firstIndex)
	fmt.Println(firstIndexString)

	lastIndex := len(name)
	lastIndexString := string(lastIndex)
	fmt.Println(lastIndex)
	fmt.Println(lastIndexString)

}
