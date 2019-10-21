package piscine

func Atoi(s string) int {
	result := 0
	/*if s[0] == '-' && s[1] == '-' || s[0] == '+' && s[1] == '+' {
		return 0
	}*/
	if s[1] < '0' || s[1] > '9' {
		return 0
	}
	for _, value := range s {
		a := 0
		if value == ' ' {
			return 0
		}
		for i := '1'; i <= value; i++ {
			a++
		}
		result = result*10 + a
	}
	if s[0] == '-' {
		result = result * -1
	}
	return result
}
