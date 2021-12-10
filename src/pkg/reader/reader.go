package reader

import (
	"fmt"
	"os"
	"path/filepath"
	"rain-csv-parser/src/pkg/reader/strategy"
	"strings"
)

type reader struct {
	extension string
	folder    string
	strategy  strategy.ReaderStrategy
}

func NewReader(folder string, extension string) (*reader, error) {
	readerStrategy, err := strategy.NewReaderStrategySelector().GetStrategy(extension)
	if err != nil {
		return nil, err
	}

	return &reader{
		extension: extension,
		folder:    folder,
		strategy:  readerStrategy,
	}, nil

}

func (r *reader) getAllFilenames(folder string) (files []string, err error) {
	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return files, err
}

func (r *reader) filterFilesByExtension(filenames []string, extension string) (csvFiles []string) {
	for _, filename := range filenames {
		fileSplit := strings.Split(filename, ".")

		if fileSplit[len(fileSplit)-1] == extension {
			csvFiles = append(csvFiles, filename)
		}
	}
	return csvFiles
}

func (r *reader) Read() error {
	filenames, err := r.getAllFilenames(r.folder)
	if err != nil {
		return err
	}
	filesToRead := r.filterFilesByExtension(filenames, r.extension)
	for _, file := range filesToRead {
		data, err := r.strategy.ReadFile(file)
		if err != nil {
			return err
		}
		fmt.Println(data)
	}
	return nil
}
