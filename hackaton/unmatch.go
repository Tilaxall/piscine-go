package piscine

func Unmatch(arr []int) int {
	len := 0
	//res := 0
	for i := range arr {
		len = i + 1
	}
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if arr[i] == arr[j] {
				arr[i], arr[j] = -1, -1
			}
		}
	}

	for _, i := range arr {
		if i != -1 {
			return i
		}
	}

	return -1
}
