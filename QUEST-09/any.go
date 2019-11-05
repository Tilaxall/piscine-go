package piscine

func Any(f func(string) bool, arr []string) bool {
	tempbool := false
	ResultBool := false
	for _, key := range arr {
		tempbool = f(key)
		if tempbool == true {
			ResultBool = true
		}
	}
	return ResultBool
}
