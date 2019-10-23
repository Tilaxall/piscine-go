package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, str := range os.Args {
		for _, tempStr := range str {
			z01.PrintRune(tempStr)
		}
		z01.PrintRune(10)
	}
}
