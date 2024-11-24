package main

import "fmt"

func main() {
	const DOMAIN string = "127.0.0.1"
	const MIN_RETRY = 1
	const MAX_RETRY = 3

	fmt.Println(DOMAIN)
	fmt.Println(MIN_RETRY)
	fmt.Println(MAX_RETRY)

	const (
		PAYMENT_METHOD          = "Mapclub Cash"
		TRANSACTION_TYPE string = "TOP UP"
	)

	fmt.Println(TRANSACTION_TYPE)
	fmt.Println(PAYMENT_METHOD)
}
