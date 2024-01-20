# CSV to JSON Formatter

## Description
This repository is a command line tool I made that takes the contents of a csv file and formats them into JSON. There are also tests that accompany the code written in this repository.

## Usage
- You will need to have Go(Golang) installed on your machine in order to run this program
- Go ahead and clone this repository and cd into this project
- Once in this project, run the following command: 'go run csv_to_json.go [path to desired csv file]'
- After this command is ran, you can navigate to the same location as the csv file and you should now see a JSON file with the same name
- The contents of this new JSON file should now have all of the csv contents formatted as JSON

## Pretty JSON
By default, the JSON this tool generates will not be pretty and it instead will be printed out on one single line. 

If you would like the JSON that is generated to be "pretty", simply run the command 'go run csv_to_json.go --pretty [path to desired csv file]'

## Notice Anything?
Please feel free to reach out or make a PR if this tool leaves anything to be desired, does not work properly for you, or you would like to see additional features