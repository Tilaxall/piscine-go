package piscine

func AppendRange(min, max int) []int {
	var ResultMass []int
	for i := min; i < max; i++ {
		ResultMass = append(ResultMass, i)
	}
	return ResultMass
}
