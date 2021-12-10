package strategy

import (
	"errors"
	"rain-csv-parser/src/domain"
	"rain-csv-parser/src/reader/strategy/csv"
)

const (
	CSVExtension = "csv"
)

type ReaderStrategy interface {
	ReadFile(filePath string) (*domain.MatrixDomain, error)
}

type readerStrategySelector struct {
	strategy ReaderStrategy
}

func NewReaderStrategySelector() *readerStrategySelector {
	return &readerStrategySelector{}
}

func (s *readerStrategySelector) GetStrategy(extension string) (ReaderStrategy, error) {
	switch extension {
	case CSVExtension:
		return csv.NewCSVReaderStrategy(), nil
	default:
		return nil, errors.New("implementation not found for this extension")
	}
}
