package main

import (
	"log"
	"os"

	"github.com/achenet/gaston/generate"
)

const _DefaultInputFile = "input.txt"

func main() {
	text, err := getTextFromInputFile()
	if err != nil {
		log.Fatalf("could not get text from input file: %s", err.Error())
	}
	transitionMatrix := generate.Generate(text)
}

func getTextFromInputFile() (string, error) {
	// read input file
	args := os.Args
	if len(args) < 2 {
		log.Printf("no input file given as argument, using %s", _DefaultInputFile)
		return read(_DefaultInputFile)
	} else {
		return read(os.Args[1])
	}

}

func read(filename string) (string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
