package piscine

func Index(s string, toFind string) int {
	findStr := []rune(toFind)
	indexStr := []rune(s)
	index := -1
	for findIndex, findKey := range findStr {
		for strIndex, key := range indexStr {
			if findKey == key && index == -1 && findIndex == 0 {
				//fmt.Println(findIndex)
				//fmt.Println(findKey)
				//fmt.Println(strIndex)
				//fmt.Println(key)
				index = strIndex
				break

			}
			if findIndex != 0 && index != -1 && findKey != key && strIndex > index {
				index = -1
			}
			if findIndex != 0 && index != -1 && findKey == key {
				break
			}

		}
	}

	/*
		findIndex := 0
		findStr := []rune(toFind)
		for keyIndex, key := range s {
			if index != 0 {
				break
			}
			if findStr[0] == key {
				index = keyIndex
				findIndex = 1
				findStr = strings.Replace
			}

		}*/

	/*for i := index; i <= len(s); i++ {
		for keyIndex, key := range s {
			if keyIndex >= index {
				if findStr[findIndex] == key {

				} else {
					return -1
				}

			}
		}
	}*/

	return index
}
