package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	len := 0
	a, b := 0, 0
	for i := range arg {
		len = i
	}
	len++
	if len == 3 {
		if IsNumb(arg[0]) && IsNumb(arg[2]) && IsOper(arg[1]) {
			a = Atoi(arg[0])
			b = Atoi(arg[2])
			fmt.Println(Solve(a, b, arg[1]))
		} else {
			fmt.Println("0")
			os.Exit(0)
		}
	} else {
		os.Exit(0)
	}

}

func Solve(a int, b int, x string) int {
	switch x {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("No division by 0")
			os.Exit(0)
		}
		return a / b
	case "%":
		if b == 0 {
			fmt.Println("No Modulo by 0")
			os.Exit(0)
		}
		return a % b
	}
	return 5
}

func IsOper(str string) bool {
	res := false
	for _, i := range str {
		if i == '+' || i == '-' || i == '*' || i == '/' || i == '%' {
			return true
		}
	}
	return res
}

func IsNumb(str string) bool {
	res := false
	for _, i := range str {
		if i >= '0' && i <= '9' {
			return true
		}
	}
	return res
}

func Atoi(s string) int {
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
