package writer

import (
	"rain-csv-parser/src/domain"
)

type (
	IOStrategy interface {
		Write(matrix [][]string, output string) error
	}
)

type writerService struct {
	ioStrategy IOStrategy
}

func NewWriterService(ioStrategy IOStrategy) *writerService {
	return &writerService{
		ioStrategy: ioStrategy,
	}
}

func (writer *writerService) Write(matrix *domain.MatrixDomain, outputPath string) error {
	return writer.ioStrategy.Write(matrix.Data, outputPath)
}
