package main

import (
	"fmt"
	"os"

	"github.com/jongedav/nins/internal/cli"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Invoke nins --help for help.")
	} else {
		request, err := cli.ParseArgs(os.Args)
		if err != nil {
			fmt.Printf("Error: %s", err)
		} else {
			fmt.Println(request)
		}

	}
}
