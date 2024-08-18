package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web req lecture LCO")
	// performGetRequest()
	// performPostJsonRequest()
	performPostFormRequest()

}
func performGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.Status)
	fmt.Println("Content length is: ", response.ContentLength)

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func performPostJsonRequest() {

	const myurl = "http://localhost:8000/post"

	// fake json playload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with GoLang",
			"Price":0,
			"Platform":"learncodeonline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

func performPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("FirstName", "Arka")
	data.Add("lastName", "Mukherjee")
	data.Add("email", "mukherjeearka445")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
