package standardizer

import "rain-csv-parser/src/domain"

type standardizerService struct {
}

func NewStandardizerService() *standardizerService {
	return &standardizerService{}
}

func (standardizerService *standardizerService) Standardize(inputMatrix *domain.MatrixDomain) (
	validOutput *domain.MatrixDomain, invalidOutput *domain.MatrixDomain, err error) {
	return
}
