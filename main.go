package main

import (
	"fmt"
	"rain-csv-parser/src/iostrategy"
	"rain-csv-parser/src/iostrategy/implementations/csv"
	"rain-csv-parser/src/parser"
	"rain-csv-parser/src/pkg/logger"
	"rain-csv-parser/src/reader"
	"rain-csv-parser/src/writer"
)

func run() error {
	logger.Info().Log("Initializing application...")
	tableColumns := createTableColumns()
	matcherSelector := createMatcherSelector()
	columnGrouper := createColumnGrouper()

	csvStrategy := csv.NewCSVStrategyImplementation()
	strategySelector := iostrategy.NewIOStrategySelector()
	strategySelector.AddStrategy(csvStrategy)
	strategy, err := strategySelector.GetStrategy(FORMAT)
	if err != nil {
		return err
	}
	readerService := reader.NewReaderService(strategy)
	writerService := writer.NewWriterService(strategy)
	parserService := parser.NewParserService(tableColumns, matcherSelector, columnGrouper)

	inputTable, err := readerService.Read(fmt.Sprintf("%s.%s", INPUT_PATH, FORMAT))
	if err != nil {
		return err
	}

	stdTable, err := parserService.Standardize(inputTable)
	if err != nil {
		return err
	}

	err = writerService.Write(stdTable, fmt.Sprintf("%s_correct.%s", OUTPUT_PATH, FORMAT))
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
