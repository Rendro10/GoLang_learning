package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to map introduction")

	languages := make(map[string]string)
	languages["JS"] = "Java Script"
	languages["PY"] = "python"
	languages["RB"] = "Ruby"

	fmt.Println("List of all languages:", languages)
	// accessing the indevidual values of map
	// fmt.Println("JS stands for: ", languages["JS"])
	// fmt.Println("PY stands for: ", languages["PY"])
	// fmt.Println("RB stands for: ", languages["RB"])

	// Delete any values from map
	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// loops are interesting go lang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
