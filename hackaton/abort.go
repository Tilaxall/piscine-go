package piscine

func Abort(a, b, c, d, e int) int {
	Mass := []int{a, b, c, d, e}
	//sort.Ints(Mass)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i-1; j++ {
			if Mass[j] > Mass[j+1] {
				temp := Mass[j]
				Mass[j] = Mass[j+1]
				Mass[j+1] = temp
			}
		}
	}

	return Mass[2]
}
