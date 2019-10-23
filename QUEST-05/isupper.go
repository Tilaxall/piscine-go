package piscine

func IsUpper(str string) bool {
	strbool := true
	for _, key := range str {
		if !(key >= 'A' && key <= 'Z') {
			strbool = false
		}
	}
	return strbool
}
