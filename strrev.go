package piscine

func PrintStr(str string) string {
	strTemp := []rune(str)
	strTempreverse := []rune(str)
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
