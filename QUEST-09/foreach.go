package piscine

func ForEach(f func(int), arr []int) {
	for _, key := range arr {
		f(key)
	}
}
