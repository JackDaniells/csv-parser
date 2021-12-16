package iostrategy

import (
	"errors"
	"fmt"
	"rain-csv-parser/src/pkg/logger"
)

type IOStrategy interface {
	CanExecute(extension string) bool
	Read(inputPath string) ([][]string, error)
	Write(matrix [][]string, outputPath string) error
}

type ioStrategySelector struct {
	strategies []IOStrategy
}

func NewIOStrategySelector() *ioStrategySelector {
	return &ioStrategySelector{
		strategies: []IOStrategy{},
	}
}

func (ioStrategy *ioStrategySelector) AddStrategy(strategy IOStrategy) {
	ioStrategy.strategies = append(ioStrategy.strategies, strategy)
}

func (ioStrategy *ioStrategySelector) GetStrategy(extension string) (IOStrategy, error) {
	for _, strategy := range ioStrategy.strategies {
		if strategy.CanExecute(extension) {
			return strategy, nil
		}
	}

	err := errors.New(fmt.Sprintf("implementation not found for %s extension", extension))
	logger.Error().Log("Error when get strategy:", err.Error())
	return nil, err
}
