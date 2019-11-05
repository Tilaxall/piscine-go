package piscine

func CountIf(f func(string) bool, tab []string) int {
	tempbool := false
	ResultBool := 0
	for _, key := range tab {
		tempbool = f(key)
		if tempbool == true {
			ResultBool++
		}
	}
	return ResultBool
}
