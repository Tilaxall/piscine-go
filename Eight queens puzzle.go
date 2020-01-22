package main

import (
	"fmt"
)

const N = 8

//var k int = 0
var position [8]int

func isSafe(int queen_number, int row_position) bool {
	for i := 0; i < queen_number; i++ {
		other_row_pos := position[i]

		if other_row_pos == row_position ||
			other_row_pos == row_position-(queen_number-i) ||
			other_row_pos == row_position+(queen_number-i) {
			return false
		}
		return true
	}
}

func solve(int k) {
	if k == N {
		fmt.Println("Решение: ")
		for i := 0; i < N; i++ {
			fmt.Println(position[i])
			fmt.Println(' ')
		}
	} else {
		for i := 0; i < N; i++ {
			if isSafe(k, i) {
				position[k] = i
				solve(k + 1)
			}
		}
	}
}

func main() {
	n := 0
	solve(n)
	return 0
}
