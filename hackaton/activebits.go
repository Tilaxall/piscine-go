package piscine

func ActiveBits(n int) uint {
	index := 0
	for n > 0 {
		if n%2 != 0 {
			index++
		}
		n = n / 2
	}
	return uint(index)
}
