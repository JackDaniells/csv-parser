package parser

import (
	"rain-csv-parser/src/commons"
	"rain-csv-parser/src/domain"
	domain2 "rain-csv-parser/src/parser/domain"
	"rain-csv-parser/src/pkg/logger"
	"strings"
)

type validatorService struct {
	tableColumns    domain.TableColumnSchemas
	matcherSelector domain2.MatchSelector
}

func NewValidatorService(tableColumns domain.TableColumnSchemas) *validatorService {
	return &validatorService{
		tableColumns: tableColumns,
	}
}

func (validatorService *validatorService) isValidElement(element string) bool {
	return strings.TrimSpace(element) != ""
}

func (validatorService *validatorService) checkAllRequiredFieldsAreFilledInRow(headers []string, row []string) bool {
	for _, col := range validatorService.tableColumns {
		if col.Required {
			colFound, colIndex := commons.FindIndexInArray(headers, col.Name)
			if !colFound || !validatorService.isValidElement(row[colIndex]) {
				return false
			}
		}
	}
	return true
}

func (validatorService *validatorService) validateAllRequiredFieldsAreFilled(inputTable *domain.TableDomain) (invalidRows []int) {
	tableHeader := inputTable.GetHeader()
	for index, row := range inputTable.GetBody() {
		if !validatorService.checkAllRequiredFieldsAreFilledInRow(tableHeader, row) {
			invalidRows = append(invalidRows, index)
		}
	}
	return invalidRows
}

//func (parserService *validatorService) verifyIfAllElementsAreUniqueInColumn(col string, matrix *domain.TableDomain) (bool, error) {
//	idColIndex := commons.FindIndexInArray(matrix.GetHeader(), col)
//	idCol, err := matrix.GetColumn(idColIndex)
//	if err != nil {
//		return false, err
//	}
//
//	idCol = commons.TrimSpacesFromArray(idCol)
//	hasDuplicatedElements := commons.HasDuplicatedElementsInArray(idCol)
//
//	return hasDuplicatedElements, nil
//}

func (validatorService *validatorService) filterTableBody(inputBody [][]string, invalidIndexes []int) (validBody [][]string, invalidBody [][]string) {
	for i, row := range inputBody {
		if commons.FindInArrayInt(invalidIndexes, i) {
			invalidBody = append(invalidBody, row)
		} else {
			validBody = append(validBody, row)
		}
	}
	return
}

func (validatorService *validatorService) Validate(inputTable *domain.TableDomain) (
	validOutput *domain.TableDomain, invalidOutput *domain.TableDomain, err error) {

	logger.Info().Log("validating if all required columns are filled...")
	reqFieldProblemRows := validatorService.validateAllRequiredFieldsAreFilled(inputTable)
	validBody, invalidBody := validatorService.filterTableBody(inputTable.GetBody(), reqFieldProblemRows)

	validOutput = domain.NewTableDomain(inputTable.GetHeader(), validBody)
	invalidOutput = domain.NewTableDomain(inputTable.GetHeader(), invalidBody)

	return

}
