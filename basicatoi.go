package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, value := range s {
		a := 0
		for i := '1'; i <= value; i++ {
			a++
		}
		result = result*10 + a
	}
}
