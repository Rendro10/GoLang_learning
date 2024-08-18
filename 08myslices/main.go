package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Concept of Slices:")
	var fruitList = []string{"Apple", "Banana", "Orange"}
	fmt.Printf("What type of data contains by fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Coconut")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 223
	highScore[1] = 210
	highScore[2] = 320
	highScore[3] = 420

	highScore = append(highScore, 550, 666, 231)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	// How to remove a value from slices based on index

	var courses = []string{"React JS", "Js", "Swift", "Python", "Rubi", "JAVA", "C++"}
	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...) // swift will be deleted using this slice method
	fmt.Println(courses)

}
