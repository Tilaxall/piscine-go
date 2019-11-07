package main

import "fmt"

func BasicAtoi2(s string) int {
	result := 0
	for _, value := range s {
		result := 0
		a := 0
		for i := '1'; i <= value; i++ {
			a++
		}
		result = result*10 + a
	}
	return result
}

func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"

	n := BasicAtoi2(s)
	n2 := BasicAtoi2(s2)
	n3 := BasicAtoi2(s3)
	n4 := BasicAtoi2(s4)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

}
