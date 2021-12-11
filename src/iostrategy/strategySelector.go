package iostrategy

import (
	"errors"
	"rain-csv-parser/src/iostrategy/implementations/csv"
)

type IOStrategy interface {
	Read(inputPath string) ([][]string, error)
	Write(matrix [][]string, outputPath string) error
}

type ioStrategySelector struct{}

func NewIOStrategySelector() *ioStrategySelector {
	return &ioStrategySelector{}
}

func (ioStrategy *ioStrategySelector) GetStrategy(strategy string) (IOStrategy, error) {
	switch strategy {
	case csv.EXTENSION:
		return csv.NewCSVImplementation(), nil
	default:
		return nil, errors.New("implementation not found for this extension")
	}
}
