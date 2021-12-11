package reader

import (
	"rain-csv-parser/src/domain"
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

func (reader *readerService) Read(inputPath string) (matrix *domain.MatrixDomain, err error) {
	data, err := reader.ioStrategy.Read(inputPath)
	if err != nil {
		return nil, err
	}

	return domain.NewMatrixDomain(data), nil

}
