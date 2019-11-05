package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	lenargs := 0
	for i := range arg {
		lenargs = i
	}
	lenargs++
	if lenargs == 1 {
		return
	} else if lenargs == 2 {
		file, err := os.Open(arg[1])
		if err != nil {
			for _, e := range string(err.Error()) {
				z01.PrintRune(e)
			}
			z01.PrintRune('\n')
		} else {
			input, _ := ioutil.ReadAll(file)
			for _, e := range string(input) {
				z01.PrintRune(e)
			}
			z01.PrintRune(10)
			file.Close()
		}
	} else if lenargs == 3 {
		file, err := os.Open(arg[1])
		file2, _ := os.Open(arg[2])
		if err != nil {
			for _, e := range string(err.Error()) {
				z01.PrintRune(e)
			}
			z01.PrintRune('\n')
		} else {

			input, _ := ioutil.ReadAll(file)
			for _, e := range string(input) {
				z01.PrintRune(e)
			}
			z01.PrintRune(10)
			input2, _ := ioutil.ReadAll(file2)
			for _, e := range string(input2) {
				z01.PrintRune(e)
			}
			file.Close()
			file2.Close()
		}
	}
}
