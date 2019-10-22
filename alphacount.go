package piscine

func AlphaCount(str string) int {
	index := 0
	for _, value := range str {
		if value >= 65 && value <= 90 || value >= 97 && value <= 122 {
			index++
		}
	}
	return index
}
