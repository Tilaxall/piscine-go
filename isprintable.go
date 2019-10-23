package piscine

func IsPrintable(str string) bool {
	strbool := true
	strRune := []rune(str)
	for _, key := range strRune {
		if key == 92 {
			strbool = false
		}
	}
	return strbool
}
