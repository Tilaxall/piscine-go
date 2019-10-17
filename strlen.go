package piscine

func StrLen(str string) int {
	strTemp := []rune(str)
	var number int
	for i := range strTemp {
		number = i
	}
	return int(number)
}
