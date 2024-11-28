package main

import (
	"fmt"
	"go-fundamental/helper"
	"go-fundamental/modifier"
	"time"
)

func main() {
	result := helper.SayHello("Muhammad Ayom Izzuddin")
	fmt.Println(result)

	modifier1 := modifier.CountTotalMessage(10, 25)
	fmt.Println(modifier1)

	application := modifier.Application
	fmt.Println(application)

	// tidak bisa diakses karena nama depan variable nya menggunakan huruf kecil
	// version := modifier.version
	// modifier2 := modifier.substractTotalMessage(1, 2)

	var uniqueCodes = []string{}
	fmt.Println(uniqueCodes)

	// Use a map to track unique codes
	uniqueCodeMap := make(map[string]struct{})

	for len(uniqueCodes) < 10 {
		uniqueCode := fmt.Sprintf("%d", time.Now().UnixNano()/1000)
		if _, exists := uniqueCodeMap[uniqueCode]; !exists {
			uniqueCodes = append(uniqueCodes, uniqueCode)
			uniqueCodeMap[uniqueCode] = struct{}{} // Mark this code as used
		}
		// Sleep for a short duration to ensure the next code is different
		time.Sleep(1 * time.Millisecond)
	}

	fmt.Println(uniqueCodes)
}
