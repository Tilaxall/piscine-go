package piscine

func FindNextPrime(nb int) int {
	for i := 1; i <= nb; i++ {
		if IsPrime(nb) == false {
			nb++
		} else {
			break
		}
	}
	return nb
}

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
