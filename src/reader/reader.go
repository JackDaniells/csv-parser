package reader

import (
	"errors"
	"rain-csv-parser/src/domain"
	"rain-csv-parser/src/pkg/logger"
)

type (
	IOStrategy interface {
		Read(input string) ([][]string, error)
	}
)

type readerService struct {
	ioStrategy IOStrategy
}

func NewReaderService(ioStrategy IOStrategy) *readerService {
	return &readerService{
		ioStrategy: ioStrategy,
	}
}

func (reader *readerService) Read(inputPath string) (table *domain.TableDomain, err error) {
	logger.Debug().Log("Reading input path...")
	data, err := reader.ioStrategy.Read(inputPath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		err = errors.New("input table is empty")
		logger.Error().Log("Error when read data from input path:", err.Error())
		return nil, err
	}

	return domain.NewTableDomain(data[0], data[1:]), nil

}
