package main

import (
	"fmt"
)

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(IterativePower(arg1, arg2))
}

func IterativePower(nb int, power int) int {
	var result int = 1
	if nb == 0 || power == 0 {
		return 0
	} else {
		for i := 1; i <= power; i++ {
			result = result * nb
		}
		return result
	}
}
