package main

import (
	"fmt"
	"os"
	"rain-csv-parser/config"
	"rain-csv-parser/src/iostrategy"
	"rain-csv-parser/src/iostrategy/implementations/csv"
	"rain-csv-parser/src/parser"
	"rain-csv-parser/src/pkg/logger"
	"rain-csv-parser/src/reader"
	"rain-csv-parser/src/validator"
	"rain-csv-parser/src/writer"
	"strings"
)

func run(filename string) error {
	logger.Info().Log("Initializing application...")
	tableColumns := config.CreateTableColumns()
	matcherSelector := config.CreateMatcherSelector()
	columnGrouper := config.CreateColumnGrouper()

	csvStrategy := csv.NewCSVStrategyImplementation()
	strategySelector := iostrategy.NewIOStrategySelector()
	strategySelector.AddStrategy(csvStrategy)
	strategy, err := strategySelector.GetStrategy(config.FORMAT)
	if err != nil {
		return err
	}

	readerService := reader.NewReaderService(strategy)
	parserService := parser.NewParserService(tableColumns, matcherSelector, columnGrouper)
	validatorService := validator.NewValidatorService(tableColumns)
	writerService := writer.NewWriterService(strategy)

	inputTable, err := readerService.Read(fmt.Sprintf("%s/%s.%s", config.INPUT_FOLDER, filename, config.FORMAT))
	if err != nil {
		return err
	}

	stdTable, err := parserService.Standardize(inputTable)
	if err != nil {
		return err
	}

	validTableOutput, invalidTableOutput := validatorService.Validate(stdTable)

	err = writerService.Write(validTableOutput, fmt.Sprintf("%s/%s_correct.%s", config.OUTPUT_FOLDER, filename, config.FORMAT))
	if err != nil {
		return err
	}

	err = writerService.Write(invalidTableOutput, fmt.Sprintf("%s/%s_bad.%s", config.OUTPUT_FOLDER, filename, config.FORMAT))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	filename := ""

	if filename == "" && len(os.Args) < 2 {
		panic("you must pass name of the file to be processed as a parameter")
	}
	filename = strings.TrimSpace(os.Args[1])
	err := run(filename)
	if err != nil {
		logger.Error().Log("Application exited with error!")
	}
}
