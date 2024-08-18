package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop Lecture")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	//type-1 for loop

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//type-2 for loop

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// type - 3 for loop

	// for index, day := range days {
	// 	fmt.Printf("index is %v and values is %v\n", index, day)
	// }

	// using continue and break in for loop

	roughValue := 1

	for roughValue < 10 {
		if roughValue == 5 {
			roughValue++
			continue
		}
		fmt.Println("value is: ", roughValue)
		roughValue++
	}

}
