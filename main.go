package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <csvFile>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	fileData, err := getFileData()

	if err != nil {
		exitGracefully(err)
	}

	if _, err := checkIfValidFile(fileData.filepath);  err != nil {
		exitGracefully(err)
	}

	writerChannel := make(chan map[string]string)
	done := make(chan bool)

	go processCsvFile(fileData, writerChannel)
	go writeJSONFile(fileData.filepath, writerChannel, done, fileData.pretty)

	<-done
 }
