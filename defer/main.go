package main

import "fmt"

func main() {
	runApplicationWithoutDefer()
	runApplicationWithDefer()
}

// function logging akan selalu dieksekusi diakhir function meskipun terjadi error pada function jika menggunakan defer

func logging() {
	fmt.Println("successFully create log.")
}

func runApplicationWithoutDefer() {
	logging()
	fmt.Println("run application without defer.")
}

func runApplicationWithDefer() {
	defer logging()
	fmt.Println("run application with defer")
}
