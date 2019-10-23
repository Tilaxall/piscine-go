package piscine

func ToLower(s string) string {
	str := []rune(s)
	for index, key := range str {
		if key >= 'A' && key <= 'Z' {
			str[index] = key + 32
		}
	}
	return string(str)
}
