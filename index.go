package main

func Index(s string, toFind string) int {
	sLen := StrLen(s)
	findLen := StrLen(toFind)
	if toFind == "" {
		return 0
	}
	var tempString string
	for i := 0; i < sLen-findLen; i++ {
		for j := 0; j < findLen; j++ {
			tempString = tempString + string(s[i+j])
		}
		if tempString == toFind {
			return i
		}
		tempString = ""
	}
	//fmt.Println(tempString)
	return -1

	/*findStr := []rune(toFind)
	indexStr := []rune(s)
	index := -1
	tempIndex := 0
	for findIndex, findKey := range findStr {
		for strIndex, key := range indexStr {
			if findKey == key && index == -1 && findIndex == 0 {
				//fmt.Println(findIndex)
				//fmt.Println(findKey)
				//fmt.Println(strIndex)
				//fmt.Println(key)
				index = strIndex
				break

			}*/

	/*if findIndex != 0 && index != -1 && findKey != key && strIndex > index+1 {
				index = -1
			}
			if findIndex != 0 && index != -1 && findKey == key && strIndex > index {
				break
			}
		}
	}*/

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

	//return index
}

func StrLen(str string) int {
	strTemp := []rune(str)
	var number int
	for i := range strTemp {
		number = i
	}
	return int(number + 1)
}
