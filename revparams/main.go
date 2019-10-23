package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lenIndex := 0
	for index, str := range os.Args {
		lenIndex = index
		str = str
	}

	for i := lenIndex; i >= 1; i-- {
		for _, key := range os.Args[i] {
			z01.PrintRune(key)
		}
		z01.PrintRune(10)
	}

}
