package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lenIndex := 0
	argument := os.Args
	for index := range os.Args {
		lenIndex = index
	}

	for i := 0; i < lenIndex-1; i++ {
		for j := 0; j < lenIndex-i-1; j++ {
			if argument[j] < argument[j+1] {
				// меняем элементы местами
				temp := argument[j]
				argument[j] = argument[j+1]
				argument[j+1] = temp
			}
		}
	}

	for i := 1; i >= lenIndex; i++ {
		for _, key := range os.Args[i] {
			z01.PrintRune(key)
		}
		z01.PrintRune(10)
	}

}
