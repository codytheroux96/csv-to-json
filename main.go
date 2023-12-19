package main

func main() {
	fileData, err := getFileData()

	if err != nil {
		exitGracefully(err)
	}
 }
