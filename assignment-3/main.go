package main

import (
	"assignment-3/controller"
	"fmt"
)

func main() {
	controller.DataBase()

	fmt.Println("load html")
	controller.LoadHTML()
}
