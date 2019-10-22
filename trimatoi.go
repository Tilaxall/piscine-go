package piscine

func TrimAtoi(s string) int {
	var resultMass []int
	var tempIndex int = -1
	var signIndex int = 0
	var numberIndex int = 0
	var sign int = -1
	var numbers int = 0
	for index, key := range s {
		if key == '-' {
			signIndex = index
		}
		if key >= '0' && key <= '9' {
			numberIndex = index
		}
		if numberIndex != 0 && signIndex != 0 {
			break
		}
	}
	for _, key := range s {
		if key >= '0' && key <= '9' {
			for i := '0'; i <= key; i++ {
				tempIndex++
			}
			resultMass = append(resultMass, tempIndex)
			tempIndex = -1
		}
	}
	for _, key := range resultMass {
		numbers = numbers*10 + key
	}
	if signIndex != 0 && numberIndex != 0 && signIndex < numberIndex {
		numbers = numbers * sign
	}
	return numbers
}
