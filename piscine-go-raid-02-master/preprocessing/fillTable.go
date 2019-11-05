package preprocessing

//PutArgsToTable puts arguments from cmd to main table
func PutArgsToTable(args *[]string, table *[9][9]int) (err bool) {
	for index, str := range *args {
		if normalizeStr(&str, args, index) {
			return true
		}
	}
	// fmt.Println(args)
	for rowIndex, rowStr := range *args {
		for colIndex, char := range rowStr {
			(*table)[rowIndex][colIndex] = int(char - '0')
		}
	}
	// fmt.Println(table)
	return false
}

//NormalizeStr converts dots to 0 in place and checks if all char are numbers
func normalizeStr(str *string, args *[]string, argsIndex int) (err bool) {
	strLen := 0
	for index := range *str {
		strLen = index + 1
	}
	if strLen != 9 {
		return true
	}

	normalized := ""
	for _, char := range *str {
		if char == '.' {
			normalized += "0"
		} else if char < '1' || char > '9' {
			return true
		} else {
			normalized += string(char)
		}
	}
	(*args)[argsIndex] = normalized
	return false
}
