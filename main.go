package main

import (
	"log"
	"os"
)

const _DefaultInputFile = "input.txt"

func main() {
	// read input file
	args := os.Args
	if len(args) < 2 {
		log.Printf("no input file given as argument, using %s", _DefaultInputFile)
	}
}
