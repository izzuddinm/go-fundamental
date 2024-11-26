package main

import "fmt"

func main() {
	startApp(false)
	startApp(true)
}

func endApp() {
	fmt.Println("End App")
}

func startApp(error bool) {
	fmt.Println("Start App")
	if error {
		panic("Oops, System unknown error.")
	}
	defer endApp()
}
