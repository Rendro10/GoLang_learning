package main

import "fmt"

func main() {
	fmt.Println("Welcome to concept of Pointer.")

	// var ptr *int
	// fmt.Println("Value of ptr is-", ptr) // default value of ptr os null

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is-", ptr)
	fmt.Println("Value of actual pointer is-", *ptr)

	*ptr = *ptr + 2

	fmt.Println("New Value is ", myNumber)
}
