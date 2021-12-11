package main

import (
	"fmt"
	"rain-csv-parser/src/iostrategy"
	"rain-csv-parser/src/reader"
	"rain-csv-parser/src/writer"
)

const (
	EXTENSION = "csv"

	OUTPUT_FILE = "output"
)

func run() error {
	fmt.Println("Initializing parser...")
	csvStrategy, err := iostrategy.NewIOStrategySelector().GetStrategy(EXTENSION)
	if err != nil {
		return err
	}
	readerService := reader.NewReaderService(csvStrategy)
	writerService := writer.NewWriterService(csvStrategy)
	fmt.Println("Parser running!")

	fmt.Println("Reading file...")
	dataMatrix, err := readerService.Read("input/roster1.csv")
	if err != nil {
		return err
	}

	fmt.Println("Writing file...")
	err = writerService.Write(dataMatrix, "output/roster1.csv")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
