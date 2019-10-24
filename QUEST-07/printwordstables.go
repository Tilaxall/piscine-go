package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, key := range table {
		for _, strKey := range key {
			z01.PrintRune(strKey)
		}
		z01.PrintRune('\n')
	}
}
