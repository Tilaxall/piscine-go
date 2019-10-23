package piscine

func Join(strs []string, sep string) string {
	var resultStr string
	strIndex := 0
	for _, key := range strs {
		resultStr = resultStr + key
		strIndex++
		if strIndex == len(strs) {
			break
		}
		resultStr = resultStr + sep

	}
	return resultStr
}
