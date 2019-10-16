package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	//var a rune = 48
	//var b rune = 49
	//var c rune = 57

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for h := '0'; h <= '9'; h++ {
				for k := h + 1; k <= '9'; k++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					//if i != 81 || j != 80 || h != 81 || k != 81 {
					z01.PrintRune(32)
					//}
					z01.PrintRune(h)
					z01.PrintRune(k)
					z01.PrintRune(44)
					z01.PrintRune(32)

				}
			}
		}
	}
	z01.PrintRune(10)
}
