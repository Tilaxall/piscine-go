package PrintStr

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	strTemp := []rune(str)
	for i := range strTemp {
		z01.PrintRune(strTemp[i])
	}
}
