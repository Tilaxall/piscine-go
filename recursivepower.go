package piscine

func RecursivePower(nb int, power int) int {
	var result int = 1
	if nb < 0 {
		return 0
	} else if nb >= 1 {
		if power > 1 {
			//nb = nb * nb
			result = result * nb * RecursivePower(nb, power-1)
		}
	} else {
		return 0
	}
	return result
}
