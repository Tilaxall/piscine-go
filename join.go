package piscine

func Join(strs []string, sep string) string {
	var resultStr string
	strIndex := 0
	massIndex := 0
	for index, key := range strs {
		key = key
		massIndex = index
	}

	for _, key := range strs {
		resultStr = resultStr + key
		strIndex++
		if strIndex == massIndex {
			break
		}
		resultStr = resultStr + sep

	}
	return resultStr
}
