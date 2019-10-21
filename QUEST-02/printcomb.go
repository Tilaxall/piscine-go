package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for h := j + 1; h <= '9'; h++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(h)
				if i != 55 || j != 56 || h != 57 {
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		}
	}
	z01.PrintRune(10)
}
