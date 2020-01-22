package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var a rune = 48
	//fmt.Println(z01.IntRange(0, 9))
	for i := a; i <= 57; i++ {
		z01.PrintRune(i)

	}
	z01.PrintRune(10)
}
