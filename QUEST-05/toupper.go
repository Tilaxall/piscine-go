package piscine

func ToUpper(s string) string {
	str := []rune(s)
	for index, key := range str {
		if key >= 'a' && key <= 'z' {
			str[index] = key - 32
		}
	}
	return string(str)
}
