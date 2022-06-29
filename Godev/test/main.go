package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var i string

	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	fmt.Println("Your string is:", i)

}
