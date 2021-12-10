package reader

import (
	"os"
	"path/filepath"
	"rain-csv-parser/src/domain"
	"rain-csv-parser/src/reader/strategy"
	"strings"
)

type readerModule struct {
	extension string
	folder    string
	strategy  strategy.ReaderStrategy
}

func NewReaderModule(folder string, extension string) (*readerModule, error) {
	readerStrategy, err := strategy.NewReaderStrategySelector().GetStrategy(extension)
	if err != nil {
		return nil, err
	}

	readerModule := &readerModule{
		extension: extension,
		folder:    folder,
		strategy:  readerStrategy,
	}

	return readerModule, nil

}

func (reader *readerModule) getAllFilenames(folder string) (files []string, err error) {
	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return files, err
}

func (reader *readerModule) filterFilesByExtension(filenames []string, extension string) (csvFiles []string) {
	for _, filename := range filenames {
		fileSplit := strings.Split(filename, ".")

		if fileSplit[len(fileSplit)-1] == extension {
			csvFiles = append(csvFiles, filename)
		}
	}
	return csvFiles
}

func (reader *readerModule) Read() (matrixes []*domain.MatrixDomain, err error) {
	filenames, err := reader.getAllFilenames(reader.folder)
	if err != nil {
		return nil, err
	}
	filesToRead := reader.filterFilesByExtension(filenames, reader.extension)

	matrixes = []*domain.MatrixDomain{}
	for _, file := range filesToRead {
		matrix, err := reader.strategy.ReadFile(file)
		if err != nil {
			return nil, err
		}
		matrixes = append(matrixes, matrix)
	}
	return matrixes, nil
}
