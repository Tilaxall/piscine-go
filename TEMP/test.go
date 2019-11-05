package main

func main() {
	stringWords := "Write a string array. (the testing will be  ours)"
	//lenght := len(stringWords)
	stringindex := 0
	//tempString := ""
	//var tempString string
	for index, key := range stringWords {
		//println(index)
		if index+1 < len(stringWords) {
			if key != ' ' && stringWords[index+1] == ' ' {
				stringindex++
			}
		}

	}
	if stringWords[len(stringWords)-1] != ' ' {
		stringindex++
	}

	for index, key := range stringWords {
		//println(index)
		if index+1 < len(stringWords) {
			if key != ' ' {

			}
		}

	}

	//println("123", len(stringWords))
	//ResultMass := make([]string, stringindex)
	//print(len(stringWords))

	println(stringindex)

}
