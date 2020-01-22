package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printResult(str string) {
	arrayStr := []rune(str)
	i := 0
	for index := range arrayStr {
		index = index
		i++
	}

	for h := 0; h < i; h++ {
		z01.PrintRune(arrayStr[h])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	//str := os.Args
	if isEven(lengthOfArg()) {
		printResult("I have an even number of arguments")
	} else {
		printResult("I have an odd number of arguments")
	}
}

func lengthOfArg() int {
	ResultIndex := 0
	for index, key := range os.Args {
		ResultIndex++
		key = key
		index = index
	}
	ResultIndex = ResultIndex - 1
	return ResultIndex
}
