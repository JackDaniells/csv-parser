package main

import (
	"fmt"
	"rain-csv-parser/config"
	"rain-csv-parser/src/iostrategy"
	"rain-csv-parser/src/iostrategy/implementations/csv"
	"rain-csv-parser/src/parser"
	"rain-csv-parser/src/pkg/logger"
	"rain-csv-parser/src/reader"
	"rain-csv-parser/src/validator"
	"rain-csv-parser/src/writer"
)

func run() error {
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

	inputTable, err := readerService.Read(fmt.Sprintf("%s.%s", config.INPUT_PATH, config.FORMAT))
	if err != nil {
		return err
	}

	stdTable, err := parserService.Standardize(inputTable)
	if err != nil {
		return err
	}

	validTableOutput, invalidTableOutput := validatorService.Validate(stdTable)

	err = writerService.Write(validTableOutput, fmt.Sprintf("%s_correct.%s", config.OUTPUT_PATH, config.FORMAT))
	if err != nil {
		return err
	}

	err = writerService.Write(invalidTableOutput, fmt.Sprintf("%s_bad.%s", config.OUTPUT_PATH, config.FORMAT))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		logger.Error().Log("Application exited with error!")
	}
}
