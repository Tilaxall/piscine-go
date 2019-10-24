package piscine

func ConcatParams(args []string) string {
	var strMass string
	for index, key := range args {
		if index == 0 {
			strMass = key
			continue
		}
		strMass += "\n" + key
	}
	return strMass
}
