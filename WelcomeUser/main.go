package main

import (
	"fmt"
	"golang-practice/utils"
)

func main() {
	fmt.Println("Please enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello, " + name)
	utils.PressEnterToExit()
}
