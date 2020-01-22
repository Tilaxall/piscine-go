package main

import (
	"fmt"
	"os"
	//"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	index := 0

	for i, key := range arguments {
		i = i
		key = key
		index++
		/* fmt.Println("index", i)
		fmt.Println("key", key) */
	}

	if index != 3 {
		fmt.Println("0")
		os.Exit(0)
	}
	//fmt.Println(index, "Test")
	number1 := BasicAtoi(os.Args[0])
	fmt.Println(number1)
	number2 := BasicAtoi(os.Args[2])
	fmt.Println(number2)
	operator := os.Args[1]
	fmt.Println(operator)

	if operator != "-" && operator != "+" && operator != "/" && operator != "*" && operator != "%" {
		fmt.Println("0")
		os.Exit(0)
	}
	//fmt.Println(index, "TEsyt")

	if index == 3 {
		switch operator {
		case "+":
			fmt.Println(number1 + number2)
		case "-":
			fmt.Println(number1 - number2)
		case "*":
			fmt.Println(number1 * number2)
		case "/":
			if number2 == 0 {
				fmt.Println("No division by 0")
				os.Exit(0)
			}
			fmt.Println(number1 / number2)
		case "%":
			if number1 == 0 {
				fmt.Println("No Modulo by 0")
				os.Exit(0)
			}
			fmt.Println(number1 % number2)
		}
	}
}

func BasicAtoi(s string) int {
	res := 0
	for _, val := range s {
		if s[0] >= '0' && s[0] <= '9' || s[0] == '+' {
			if val == 32 {
				return 0
			}
			a := 0
			for i := '1'; i <= val; i++ {
				a++
			}
			res = res*10 + a
		} else if s[0] == '-' {
			a := 0
			for i := '1'; i <= val; i++ {
				a--
			}
			res = res*10 + a
		}
	}
	return res
}
