package piscine

func Fibonacci(index int) int {
	if index < 2 {
		return index
	} else {
		return (Fibonacci(index-1) + Fibonacci(index-2))
	}
}
