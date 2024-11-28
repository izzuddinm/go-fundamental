package main

import (
	"fmt"
	"go-fundamental/helper"
	"go-fundamental/modifier"
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
}
