package main

import (
	"fmt"
)

func main() {
	str := "     Hello how are you?"
	fmt.Println(SplitWhiteSpaces1(str))
}

func SplitWhiteSpaces1(str string) []string {
	Index := 0
	indexResultMass := 0
	len := 0
	for _, slovo := range str {
		len++

	}
	massiv := make([]string, len)
	var res []string
	for i, slovo := range massiv {
		res[i] = string(slovo)

	}
	return res[]
}
