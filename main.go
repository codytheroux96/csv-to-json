package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
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

func checkIfValidFile(filename string) (bool, error) {
	if fileExtenstion := filepath.Ext(filename); fileExtenstion != ".csv" {
		return false, fmt.Errorf("File %s is not a CSV file", filename)
	}
	 if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("File %s does not exist!", filename)
	 }
	 return true, nil
}

func main() {
	fileData, err := getFileData()

	if err != nil {
		exitGracefully(err)
	}
}
