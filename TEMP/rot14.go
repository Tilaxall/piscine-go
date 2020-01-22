package piscine

/* package main

import (
	"github.com/01-edu/z01"
)

func main() {
	result := Rot14("ypKnqR30mKe7j")
	arrayRune := []rune(result)

	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
} */

func Rot14(str string) string {
	strT := []rune(str)
	/* strF := (str)
	FindBool := false
	//index := 13
	for _, key := range strT {
		if key >= 0 && key <= 9 {
			FindBool = true

		}
	} */
	for index, key := range strT {
		index = index
		if (key < 91 && key > 76) || (key < 123 && key > 108) {
			key = key - 12
		} else if (key < 77 && key > 64) || (key < 109 && key > 96) {
			key = key + 14
		}
	}

	/* if FindBool == false {
		return string(str)
	}
	return string(strF) */

	return string(str)

}

/* package student

func Rot14(str string) string {
	runes := []rune(str)
	for _, i := range runes {
		if i >= 'a' && i <= 'z' {
			if i+14 <= 122 {
				i += 14
			} else {
				i = i - 26 + 14
			}
		} else if i >= 'A' && i <= 'Z' {
			if i+14 <= 90 {
				i += 14
			} else {
				i = i - 26 + 14
			}
		}
	}
	rot14 := string(runes)
	return rot14
}
*/
