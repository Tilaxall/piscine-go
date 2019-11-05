package main

import (
	"fmt"
	"os"
)

func main() {
	//findBool := false
	for _, key := range os.Args {
		if key == "01" || key == "galaxy" || key == "galaxy 01" {
			//findBool = true
			fmt.Println("Alert!!!")
			break
		}

		/* if findBool == true {

		} */
	}
}
