package piscine

func IsPrintable(str string) bool {
	strbool := true
	strRune := []rune(str)
	for _, key := range strRune {
		if key >= 0 && key <= 31 {
			strbool = false
		}
	}
	return strbool
}
