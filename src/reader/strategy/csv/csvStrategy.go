package csv

import (
	"encoding/csv"
	"os"
	"rain-csv-parser/src/domain"
)

type csvReaderStrategy struct {
}

func NewCSVReaderStrategy() *csvReaderStrategy {
	return &csvReaderStrategy{}
}

func (csvStrategy *csvReaderStrategy) ReadFile(filePath string) (matrix *domain.MatrixDomain, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	stringMatrix, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	matrix = &domain.MatrixDomain{
		Data: stringMatrix,
	}
	return matrix, nil
}
