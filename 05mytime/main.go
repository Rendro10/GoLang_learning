package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study on GOLang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createrDate := time.Date(2001, time.April, 17, 23, 23, 0, 0, time.UTC)
	fmt.Println(createrDate)
	fmt.Println(createrDate.Format("01/02/2006 Monday"))

}
