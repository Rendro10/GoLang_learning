package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is all about user Input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter The rating of our food:-")

	// comma ok syntax or error ok syntax if any error comes with user input then it stores into _ otherwise it store the value in input

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating ", input)
}
