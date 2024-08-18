package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("If Else in GoLang")
	logiCount := 23
	var result string
	if logiCount < 10 {
		result = "Regular User"
	} else {
		result = "Not Regular User"
	}
	fmt.Println(result)

	//if else statement with the help of user input using bufio

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Age:-")

	input, _ := reader.ReadString('\n')
	var ans string
	if input < "18" {
		ans = "You cant drive"
	} else {
		ans = "You can drive"
	}
	fmt.Println(ans)
}
