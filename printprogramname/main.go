package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	byteStr := []rune(os.Args[0])
	for _, str := range byteStr {
		z01.PrintRune(str)
	}
	z01.PrintRune(10)
}
