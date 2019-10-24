package piscine

func SplitWhiteSpaces(str string) []string {
	Index := 0
	indexResultMass := 0
	for index, key := range str {
		if key == ' ' && !(key >= 0 && key <= 31) {
			indexResultMass++
		}
		Index = index
	}

	var tempStr string
	resultMass := make([]string, indexResultMass-2)
	for keyIndex, key := range str {
		if !(key >= 0 && key <= 31) {
			tempStr += string(key)
		}
		if key == ' ' || keyIndex == Index {
			for i := 0; i <= Index; i++ {
				resultMass[i] = tempStr
				break
			}
		}
	}
	return resultMass
}
