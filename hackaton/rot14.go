package piscine

func Rot14(str string) string {
	runes := []rune(str)
	for i := range runes {
		if runes[i] >= 97 && runes[i] <= 122 {
			if runes[i]+14 <= 122 {
				runes[i] = runes[i] + 14
			} else {
				runes[i] = runes[i] - 26 + 14
			}
		} else if runes[i] >= 65 && runes[i] <= 90 {
			if runes[i]+14 <= 90 {
				runes[i] = runes[i] + 14
			} else {
				runes[i] = runes[i] - 26 + 14
			}
		}
	}
	rot14 := string(runes)
	return rot14
}
