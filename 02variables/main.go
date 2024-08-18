package main

import "fmt"

func main() {
	// fmt.Println("Variables")

	var username string = "Arka"
	fmt.Println(username)
	fmt.Printf("variable of the type of->%T \n", username)

	var isLogges bool = false
	fmt.Println(isLogges)
	fmt.Printf("variable of the type of->%T \n", isLogges)

	var smallValue uint8 = 255
	// var smallValue uint8 = 256 this will give an error coz unit8 range is 255
	fmt.Println(smallValue)
	fmt.Printf("variable of the type of->%T \n", smallValue)

	var defaultVal int
	fmt.Println(defaultVal) //it not assign any garbeg
	fmt.Printf("variable of the type of->%T \n", defaultVal)

	// string storing

	var website = "Arka's World"
	fmt.Println(website)

	// another type of declearation with := i can assign value without datatpe but inside any method you are allowed to do this i can't do this outside the main function

	numberOfUser := 2000
	fmt.Println(numberOfUser)
	fmt.Printf("variable of the type of->%T \n", numberOfUser)

}
