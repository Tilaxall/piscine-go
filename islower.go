package piscine

func IsLower(str string) bool {
	strbool := true
	for _, key := range str {
		if !(key >= 'a' && key <= 'z') {
			strbool = false
		}
	}
	return strbool
}
