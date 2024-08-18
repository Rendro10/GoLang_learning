package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"CourseName"`
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"_"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	lcocourses := []course{
		{"ReactJS Bootcamp", 299, "Learncodeonline.in", "abc123", []string{"Web-dev", "JS"}},
		{"Mern Bootcamp", 199, "Learncodeonline.in", "bcs123", []string{"Full Stack", "JS"}},
		{"Angular Bootcamp", 299, "Learncodeonline.in", "lsd123", nil},
	}
	finalJson, err := json.MarshalIndent(lcocourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func main() {
	fmt.Println("Welcome to create json data in Golang lecture")
	// EncodeJson()
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"CourseName": "ReactJS Bootcamp",
			"Price": 299,
			"Website": "Learncodeonline.in", 
			"_": "abc123",
			"tags": ["Web-dev","JS"]
		}
	`)

	var lcocourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("Json is not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
}
