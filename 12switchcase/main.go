package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Swith case in GoLang with example of Dice game")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is :-", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("The Dice value is 1 and You can open")
	case 2:
		fmt.Println("You can move to 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("The Dice value is 6 and You can roll a new Dice")

	default:
		fmt.Println("What was that?")

	}
}
