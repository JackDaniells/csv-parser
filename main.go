package main

import (
	"fmt"
	"rain-csv-parser/src/commons/constants"
	"rain-csv-parser/src/reader"
)

const (
	INPUT_FOLDER = "input"
	OUTPUT_FILE  = "output"
)

func run() error {
	fmt.Println("Initializing parser...")

	readerCSV, err := reader.NewReaderModule(INPUT_FOLDER, constants.CSV_EXTENSION)
	if err != nil {
		return err
	}

	fmt.Println("Parser running!")

	fmt.Println(fmt.Sprintf("Reading all files from %s folder", INPUT_FOLDER))
	files, err := readerCSV.Read()
	if err != nil {
		return err
	}

	fmt.Println(files)

	return nil

}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
