package main

import (
	"rain-csv-parser/src/iostrategy"
	"rain-csv-parser/src/parser"
	"rain-csv-parser/src/pkg/logger"
	"rain-csv-parser/src/reader"
	"rain-csv-parser/src/writer"
)

const (
	EXTENSION = "csv"
)

func run() error {
	logger.Info().Log("Initializing parser...")
	csvStrategy, err := iostrategy.NewIOStrategySelector().GetStrategy(EXTENSION)
	if err != nil {
		logger.Error().Log("Error when create strategy:", err.Error())
		return err
	}
	readerService := reader.NewReaderService(csvStrategy)
	writerService := writer.NewWriterService(csvStrategy)
	parserService := parser.NewParserService()
	logger.Info().Log("Parser running!")

	logger.Info().Log("Reading input path...")
	inputMatrix, err := readerService.Read("input/roster.csv")
	if err != nil {
		logger.Error().Log("Error when read data from input path:", err.Error())
		return err
	}
	logger.Info().Log("Read input completed!")

	logger.Info().Log("Parsing input data...")
	correctDataMatrix, badDataMatrix, err := parserService.Standardize(inputMatrix)
	if err != nil {
		logger.Error().Log("Error when parse input data:", err.Error())
	}
	logger.Info().Log("Parsing data completed!")

	logger.Info().Log("Writing output with correct data...")
	err = writerService.Write(correctDataMatrix, "output/roster1_correct.csv")
	if err != nil {
		logger.Error().Log("Error when write correct output data:", err.Error())
		return err
	}
	logger.Info().Log("Write output correct data completed!")

	logger.Info().Log("Writing output with bad data...")
	err = writerService.Write(badDataMatrix, "output/roster1_bad.csv")
	if err != nil {
		logger.Error().Log("Error when write bad output data:", err.Error())
		return err
	}
	logger.Info().Log("Write output bad data completed!")

	return nil
}

func main() {
	err := run()
	if err != nil {
		logger.Error().Log("Application exited with error!")
	}
}
