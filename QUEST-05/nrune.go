package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}

func NRune(s string, n int) rune {
	//runeStr := []rune(s)
	//var number int
	byteStr := []rune(s)
	return byteStr[n-1]
	//return runeStr(-3:s[len(s)])
	/*for i := range runeStr {
		number = i
	}
	if n > number {

	} else {
		return runeStr[n+1]
	}
	return ' '*/
}
