package piscine

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
