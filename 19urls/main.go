package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://courses.learncodeonline.in/learn"

func main() {
	fmt.Println("Welcome to handling urls in GoLang")
	fmt.Println(myurl)

	//parse
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query params are:- %T\n", qparams)

	fmt.Println(qparams["coursenmae"])
}
