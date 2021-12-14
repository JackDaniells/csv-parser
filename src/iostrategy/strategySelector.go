package iostrategy

import (
	"errors"
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
	return nil, errors.New("implementation not found for this extension")
}
