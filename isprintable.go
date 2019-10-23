package piscine

func IsPrintable(str string) bool {
	strbool := true
	strRune := []rune(str)
	for _, key := range strRune {
		if key == '\\' {
			strbool = false
		}
	}
	return strbool
}
