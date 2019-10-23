package piscine

func BasicJoin(strs []string) string {
	var resultStr string
	for _, key := range strs {
		resultStr = resultStr + key
	}
	return resultStr
}
