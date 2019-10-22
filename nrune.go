package piscine

func NRune(s string, n int) rune {
	//runeStr := []rune(s)
	var number = 0
	byteStr := []rune(s)
	for _, key := range s {
		key = key
		number++
	}
	if number < n {
		return '\x00'
	} else if n <= 0 {
		return '\x00'
	}
	return byteStr[n-1]
	//return runeStr(-3:s[len(s)])
	/*for i := range runeStr {
		number = i
	}
	if n > number {

	} else {
		return runeStr[n+1]
	}
	return ' '*/
}
