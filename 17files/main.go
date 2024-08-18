package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the concept of files handling")

	content := "This needs to go in a file -LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is:-", length)

	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data into the file is ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
