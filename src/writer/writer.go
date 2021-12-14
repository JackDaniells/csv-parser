package writer

import (
	"fmt"
	"rain-csv-parser/src/domain"
	"rain-csv-parser/src/pkg/logger"
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

func (writer *writerService) Write(table *domain.TableDomain, outputPath string) error {
	logger.Debug().Log(fmt.Sprintf("Writing output in %s...", outputPath))

	outputMatrix := [][]string{
		table.GetHeader(),
	}
	outputMatrix = append(outputMatrix, table.GetBody()...)

	err := writer.ioStrategy.Write(outputMatrix, outputPath)
	if err != nil {
		logger.Error().Log("Error when write output data: ", err.Error())
		return err
	}

	return nil
}
