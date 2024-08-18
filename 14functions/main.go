package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in GoLang")
	greater()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	sumRes := sum(2, 4, 5, 6, 7)
	fmt.Println("Result of sum is: ", sumRes)
}

// here int is the return type of the function
func adder(val1 int, val2 int) int {
	return (val1 + val2)
}

func sum(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
func greater() {
	fmt.Println("Namastey from goLang")
}
