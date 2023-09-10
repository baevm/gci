package main

import (
	"fmt"
	"gci/repl"
	"os"
)

func main() {
	fmt.Printf("Welcome to GCI v1.0.0. \n")

	repl.Start(os.Stdin, os.Stdout)
}
