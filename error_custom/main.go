package main

import (
	"fmt"
)

type ValidationError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"validation error!"}
	} else if id != "izzuddin" {
		return &NotFoundError{"data not found!"}
	}
	return nil
}

func main() {
	err := SaveData("IDGOLANG", `{"name":"Muhammad Ayom Izzuddin"}`)
	if err != nil {
		// fmt.Println(err)
		// if validationErr, ok := err.(*ValidationError); ok {
		// 	fmt.Println("validation error! :", validationErr.Message)
		// } else if notFoundErr, ok := err.(*NotFoundError); ok {
		// 	fmt.Println("not found error! :", notFoundErr.Message)
		// } else {
		// 	fmt.Println("unknown error!")
		// }

		switch finalErr := err.(type) {
		case *ValidationError:
			fmt.Println("validation error! :", finalErr.Message)
		case *NotFoundError:
			fmt.Println("not found error! :", finalErr.Message)
		default:
			fmt.Println("unknown error!")
		}
	} else {
		fmt.Println("sukses")
	}
}
