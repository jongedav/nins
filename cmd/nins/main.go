package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invoke nins --help for help.")
	} else {
		switch os.Args[1] {
		case "status":
			fmt.Println("NINS Status.")
		default:
			fmt.Println("Unkown argument. Invoke nins --help for help.")
		}
	}

}
