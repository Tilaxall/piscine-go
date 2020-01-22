package main

import (
	"fmt"
	"os"
)

func main() {
	//must succeed
	//args := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	var args [9]string
	for index, key := range os.Args {
		if index != 0 {
			args[index-1] = key
		}
	}

	//must fail
	//args := []string{".96.4...1", "1...6.1.4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}

	table := [9][9]int{}
	//preprocessing functions are in 'preprocessing' folder
	if PutArgsToTable(&args, &table) {
		fmt.Println("Error")
		return
	}
	//fmt.Println(args)

	//fmt.Println(table)

	// tableCopy := table
	// tableCopy[0] = table[0]

	if backtrack(&table) {
		fmt.Println("The Sudoku was solved successfully:")
		printBoard(table)
	} else {
		fmt.Printf("The Sudoku can't be solved.")
	}
	// fmt.Println(table, tableCopy)
}

func backtrack(board *[9][9]int) (err bool) {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func printBoard(board [9][9]int) {
	//fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			//fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

//PutArgsToTable puts arguments from cmd to main table
func PutArgsToTable(args *[9]string, table *[9][9]int) (err bool) {
	for index, str := range *args {
		if normalizeStr(&str, args, index) {
			return true
		}
	}
	// fmt.Println(args)
	for rowIndex, rowStr := range *args {
		for colIndex, char := range rowStr {
			(*table)[rowIndex][colIndex] = int(char - '0')
		}
	}
	// fmt.Println(table)
	return false
}

//NormalizeStr converts dots to 0 in place and checks if all char are numbers
func normalizeStr(str *string, args *[9]string, argsIndex int) (err bool) {
	strLen := 0
	for index := range *str {
		strLen = index + 1
	}
	if strLen != 9 {
		return true
	}

	normalized := ""
	for _, char := range *str {
		if char == '.' {
			normalized += "0"
		} else if char < '1' || char > '9' {
			return true
		} else {
			normalized += string(char)
		}
	}
	(*args)[argsIndex] = normalized
	return false
}
