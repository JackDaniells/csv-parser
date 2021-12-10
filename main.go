package main

import (
	"fmt"
	reader "rain-csv-parser/src/pkg/reader"
)

const (
	INPUT_FOLDER = "input"
	OUTPUT_FILE  = "output"

	CSV_EXTENSION = "csv"
)

func run() {
	fmt.Println("Initializing parser...")

	readerModule, err := reader.NewReader(INPUT_FOLDER, CSV_EXTENSION)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Parser running!")

	fmt.Println(fmt.Sprintf("Reading all files from %s folder", INPUT_FOLDER))
	readerModule.Read()

}

func main() {
	run()
}
