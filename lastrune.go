package piscine

func LastRune(s string) rune {
	byteStr := []rune(s)
	number := 0
	for _, str := range s {
		str = str
		number++
	}
	return byteStr[number-1]
}
