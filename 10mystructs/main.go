package main

import "fmt"

func main() {
	fmt.Println("Struct in Go Lang")

	// No inheritance in go lang no super or no parent
	arka := User{"Arka", "mukherjee@golang", true, 22}
	fmt.Println(arka)
	// i will get details along with the lable here in organized way
	fmt.Printf("Arka's Details are: %+v\n", arka)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
