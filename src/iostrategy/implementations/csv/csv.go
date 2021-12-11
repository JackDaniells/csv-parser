package csv

import (
	"encoding/csv"
	"io"
	"os"
)

const (
	EXTENSION = "csv"
)

type (
	OSManager interface {
		Open(name string) (*os.File, error)
		Create(name string) (*os.File, error)
	}
)

type csvImplementation struct{}

func NewCSVImplementation() *csvImplementation {
	return &csvImplementation{}
}

func (csvImp *csvImplementation) openFile(inputPath string) (io.ReadCloser, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (csvImp *csvImplementation) createFile(outputPath string) (io.WriteCloser, error) {
	file, err := os.Create(outputPath)
	if err != nil {
		return nil, err
	}
	return file, err
}

func (csvImp *csvImplementation) readCSVFile(reader io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(reader)
	return csvReader.ReadAll()
}

func (csvImp *csvImplementation) writeCSVFile(file io.Writer, matrix [][]string) error {
	csvWriter := csv.NewWriter(file)
	for _, row := range matrix {
		err := csvWriter.Write(row)
		if err != nil {
			return err
		}
	}
	defer csvWriter.Flush()
	return nil
}

func (csvImp *csvImplementation) Read(inputPath string) ([][]string, error) {
	file, err := csvImp.openFile(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return csvImp.readCSVFile(file)
}

func (csvImp *csvImplementation) Write(matrix [][]string, outputPath string) error {
	file, err := csvImp.createFile(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return csvImp.writeCSVFile(file, matrix)
}
