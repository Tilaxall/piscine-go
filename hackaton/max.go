package puscine

func Max(arr []int) int {
	max := 0
	for index, key := range arr {
		if index == 0 {
			max = key
		} else if key > max {
			max = key
		}
	}
	return max
}
