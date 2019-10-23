package piscine

func IsNumeric(str string) bool {
	strbool := true
	for _, key := range str {
		if !(key >= '0' && key <= '9') {
			strbool = false
		}
	}
	return strbool
}
