package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lenIndex := 0
	letterBool := false
	letterIndex := 96
	//argument := os.Args
	for index, key := range os.Args {
		lenIndex = index
		if key == "--upper" {
			letterBool = true
		}
	}
	if letterBool == true {
		letterIndex = 64
	}
	for i := 1; i <= lenIndex; i++ {
		for _, key := range os.Args[i] {
			if key != "--upper" {
				z01.PrintRune(rune(key + letterIndex))
			}
		}
		z01.PrintRune(' ')
	}

}
