package main

import "fmt"

func main() {
	// defer helps to cut out the line and add to the last like here world will print after hello even after run before than hello tough
	defer fmt.Println("World!")
	fmt.Println("Hello")
	myDefer()

}

// defer works like a stack here I marked differ before fmt.println() so 0,1,2,3,4 will stored into a stack and then it will print so sequence of printing will be 4,3,2,1,0.
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
