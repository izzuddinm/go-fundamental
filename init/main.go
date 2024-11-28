package main

import (
	"fmt"
	"go-fundamental/database"
	_ "go-fundamental/internal"
)

func main() {
	db := database.GetDatabase()
	fmt.Println(db)
}
