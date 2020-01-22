package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	res := true
	len := 0
	for i := range tab {
		len = i
	}

	for i := range tab {
		if i == len {
			return true
		} else if f(tab[i], tab[i+1]) >= 0 {
			if i == len-1 {
				continue
			} else if f(tab[i+1], tab[i+2]) < 0 {
				return false
			}
			continue
		} else if f(tab[i], tab[i+1]) <= 0 {
			if i == len-1 {
				continue
			} else if f(tab[i+1], tab[i+2]) > 0 {
				return false
			}
			continue
		}
	}
	return res
}
