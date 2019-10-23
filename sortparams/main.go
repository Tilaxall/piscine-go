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

	for i := 1; i <= lenIndex; i++ {
		for j := i + 1; j <= lenIndex; j++ {
			if argument[i] > argument[j] {
				argument[i], argument[j] = argument[j], argument[i]
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
