package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var fileName string
	lenIndex := 0
	for index, str := range os.Args {
		lenIndex = index
		fileName = str
	}
	lenIndex = lenIndex + 1
	if lenIndex > 2 {
		fmt.Println("Too many arguments")
	} else if lenIndex < 1 || fileName != "quest8.txt" {
		fmt.Println("File name missing")
	} else {
		//file, err := os.Open("quest8.txt")
		fileTest, err := ioutil.ReadFile("quest8.txt")

		if err != nil {
			fmt.Printf(err.Error())
		}
		fmt.Println(string(fileTest))
	}
}
