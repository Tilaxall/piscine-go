package piscine

func IterativePower(nb int, power int) int {
	var result int = 0
	var res int = 1
	if nb <= 0 || power <= 0 {
		return 0
	} else {
		for i := 1; i <= power; i++ {
			res = res * nb
		}
		return res
	}
}
