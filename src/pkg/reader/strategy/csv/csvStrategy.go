package csv

import (
	"encoding/csv"
	"os"
	"rain-csv-parser/src/pkg/reader/domain"
)

type csvReaderStrategy struct {
}

func NewCSVReaderStrategy() *csvReaderStrategy {
	return &csvReaderStrategy{}
}

func (c *csvReaderStrategy) parseCSVToInputMatrix(csvData [][]string) (inputMatrix *domain.InputMatrix) {
	inputMatrix = &domain.InputMatrix{}
	for _, csvRow := range csvData {
		row := domain.Row{Columns: csvRow}
		inputMatrix.AddRow(row)
	}
	return inputMatrix
}

func (c *csvReaderStrategy) ReadFile(filePath string) (*domain.InputMatrix, error) {
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
	return c.parseCSVToInputMatrix(stringMatrix), nil
}
