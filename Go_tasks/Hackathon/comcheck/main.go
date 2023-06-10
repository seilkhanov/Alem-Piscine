package main

import (
	"fmt"
	"os"
)

func main() {
	check := 0
	for _, arg := range os.Args[1:] {
		alert := arg == "01" || arg == "galaxy" || arg == "galaxy 01"
		if alert {
			check++
		}
	}
	if check != 0 {
		fmt.Println("Alert!!!")
	}
}
