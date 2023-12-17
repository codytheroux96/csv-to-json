package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type inputFile struct {
	filepath  string
	separator string
	pretty    bool
}

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func getFileData() (inputFile, error) {
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("Must provide a filepath argument")
	}

	separator := flag.String("separator", "comma", "Column separator")
	pretty := flag.Bool("pretty", false, "Generate pretty JSON")

	flag.Parse()

	fileLocation := flag.Arg(0)

	if !(*separator == "comma" || *separator == "semicolon") {
		return inputFile{}, errors.New("Can only use comma or semicolon separators")
	}

	return inputFile{fileLocation, *separator, *pretty}, nil

}

func main() {
	fileData, err := getFileData()

	if err != nil {
		exitGracefully(err)
	}
}
