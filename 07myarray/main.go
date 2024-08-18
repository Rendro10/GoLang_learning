package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array concept in GOLANG")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Pine-Apple"
	fruits[2] = "Banana"
	fruits[3] = "Orange"

	fmt.Println("Fruit list is :- ", fruits)
	fmt.Println("length of the Fruit list is :- ", len(fruits))
}
