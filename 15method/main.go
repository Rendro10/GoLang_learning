package main

import "fmt"

func main() {
	fmt.Println("Struct in Go Lang")

	// No inheritance in go lang no super or no parent
	arka := User{"Arka", "mukherjee@golang", true, 22}
	fmt.Println(arka)
	// i will get details along with the lable here in organized way
	fmt.Printf("Arka's Details are: %+v\n", arka)
	arka.GetStatus()
	arka.newMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is Useer Active. ", u.Status)
}

func (u User) newMail() {
	u.Email = "test@go.gmail"
	fmt.Println("Email of the user is:-", u.Email)
}
