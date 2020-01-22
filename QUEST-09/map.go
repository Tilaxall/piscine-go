package piscines

func Map(f func(int) bool, arr []int) []bool {
	index := 0
	for tempIndex := range arr {
		index++
		tempIndex = tempIndex
	}

	ResultMass := make([]bool, index)

	for tempindex, key := range arr {
		ResultMass[tempindex] = f(key)
	}
	return ResultMass
}
