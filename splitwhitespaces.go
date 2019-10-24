package piscine

// SplitWhiteSpaces function separates the words of a string
// and puts them in a string array.
// The separators are spaces, tabs and newlines.
func SplitWhiteSpaces(str string) []string {
	size := 1
	var result []string
	lenstr := 0
	for i := range str {
		lenstr = i + 1
	}

	for i := 0; i < lenstr-1; i++ {
		if str[i] == ' ' || str[i] == '\t' || str[i] == '\n' {
			size++
			if i > 0 {
				if str[i-1] == ' ' || str[i-1] == '\t' || str[i-1] == '\n' {
					size--
				}
			}
		}
	}

	result = make([]string, size)

	tempstr := ""
	j := 0
	for i := 0; i <= lenstr; i++ {
		if i == lenstr {
			if tempstr != "" {
				result[j] = tempstr
			}
		} else if str[i] != ' ' && str[i] != '\t' && str[i] != '\n' {
			tempstr = tempstr + string(str[i])
		} else {
			if tempstr != "" {
				result[j] = tempstr
				j++
			}
			tempstr = ""
		}
	}
	return result
}
