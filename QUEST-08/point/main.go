package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func toRune(num int) {
	if num < 5 {
		for i, j := 0, '0'; i <= num%10; i, j = i+1, j+1 {
			if i == num%10 {
				z01.PrintRune(j)
			}
		}
	} else {
		nextnum := num / 10
		mod := num % 10
		toRune(nextnum)
		toRune(mod)
	}

}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	toRune(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	toRune(points.y)
	z01.PrintRune('\n')

	//fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
