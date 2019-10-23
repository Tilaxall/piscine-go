package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for index, str := range os.Args {
		for _, tempStr := range str {
			if index == 0 {
				break
			}
			z01.PrintRune(tempStr)
		}
		if index == 0 {
		} else {
			z01.PrintRune(10)
		}

	}
}
