package main

import (
	"acg/services"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {

		args := os.Args[1:]

		switch args[0] {

		// input args[1]; output args[2];
		case "--parse":
			services.ParseAlphabet(args[1], args[2])

		case "--generate":
			services.GenerateCard()

		}
	} else {
		fmt.Println("Wake up Neo")
	}
}