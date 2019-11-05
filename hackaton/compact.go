package piscine

func Compact(ptr *[]string) int {
	index := 0
	for _, key := range *ptr {
		if key != "" {
			index++
		}
	}

	temp := make([]string, index)
	count := 0
	for _, str := range *ptr {
		if str != "" {
			temp[count] = str
			count++
		}
	}
	*ptr = temp
	//ResultArr := make([]string, index)
	return count
}
