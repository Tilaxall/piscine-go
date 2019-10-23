package piscine

func IsPrintable(str string) bool {
	strbool := true
	for _, key := range str {
		if !(key >= 'a' && key <= 'z' || key >= 'A' && key <= 'Z' || key >= '0' && key <= '9') {
			strbool = false
		}
	}
	return strbool
}
