package main

import (
	"fmt"
)

func main() {
	startApp(false)
	startApp(true)
}

// contoh penggunaan recover() yang salah

/*
func endApp() {
	fmt.Println("End App")
}

func startApp(error bool) {
	fmt.Println("Start App")
	defer endApp()
	if error {
		panic("Oops, System unknown error.")
	}
}
*/

// panic harus digunakan didalam defer

// contoh penggunaan recov er() yang benar
func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Happen a panic()", message)
}

func startApp(error bool) {
	defer endApp()
	if error {
		panic("Oops, System unknown error.")
	}
}
