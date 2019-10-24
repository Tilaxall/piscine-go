package piscine

func MakeRange(min, max int) []int {

	if min >= max {
		return nil
	}

	ResultMass := make([]int, max-min)
	index := 0
	for key := range ResultMass {
		ResultMass[key] = min + index
		index++
	}

	/*for i := min; i < max; i++ {
		for index := range ResultMass {
			ResultMass[index] = i
			break
		}
	}*/
	return ResultMass
}
