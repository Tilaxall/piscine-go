package piscine

func StrRev(s string) string {
	strTemp := []rune(s)
	strTempreverse := []rune(s)
	var number int
	for i := range strTemp {
		number = i
	}
	for i := range strTemp {
		strTempreverse[i] = strTemp[number]
		number--
	}
	return string(strTempreverse)
}
