package validator

import (
	"errors"
	"rain-csv-parser/src/commons"
	"rain-csv-parser/src/domain"
	"rain-csv-parser/src/pkg/logger"
	"strings"
)

type validatorService struct {
	tableColumns domain.TableColumnSchemas
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

func (validatorService *validatorService) validateRequiredFieldsAreFilled(inputTable *domain.TableDomain) (invalidRows []int) {
	tableHeader := inputTable.GetHeader()
	for index, row := range inputTable.GetBody() {
		if !validatorService.checkAllRequiredFieldsAreFilledInRow(tableHeader, row) {
			invalidRows = append(invalidRows, index)
		}
	}
	return invalidRows
}

func (validatorService *validatorService) checkDuplicatedFieldsForUniqueColumn(table *domain.TableDomain, headerName string) (
	duplicatedElements [][]int, err error) {
	headerFound, idColIndex := commons.FindIndexInArray(table.GetHeader(), headerName)
	if !headerFound {
		return nil, errors.New("unique header name not found")
	}
	colData, err := table.GetColumn(idColIndex)
	if err != nil {
		return nil, err
	}
	colData = commons.TrimSpacesFromArray(colData)
	duplicatedElements = commons.GetDuplicatedElementsIndexesInArray(colData)
	return duplicatedElements, nil
}

func (validatorService *validatorService) validateUniqueElements(table *domain.TableDomain) (invalidIndexes []int) {
	for _, col := range validatorService.tableColumns {
		if col.Unique {
			duplicateds, err := validatorService.checkDuplicatedFieldsForUniqueColumn(table, col.Name)
			if err != nil {
				for index := range table.GetBody() {
					invalidIndexes = append(invalidIndexes, index)
				}
				return invalidIndexes
			}
			invalidIndexes = append(invalidIndexes, commons.ConvertMatrixToArray(duplicateds)...)
		}
	}
	return commons.RemoveDuplicatedFields(invalidIndexes)
}

func (validatorService *validatorService) filterInvalidIndexesInTableBody(inputBody [][]string, invalidIndexes []int) (validBody [][]string, invalidBody [][]string) {
	validBody = [][]string{}
	invalidBody = [][]string{}
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
	validOutput *domain.TableDomain, invalidOutput *domain.TableDomain) {

	logger.Info().Log("validating if all required columns are filled...")
	reqFieldProblemRows := validatorService.validateRequiredFieldsAreFilled(inputTable)
	reqValidBody, reqInvalidBody := validatorService.filterInvalidIndexesInTableBody(inputTable.GetBody(), reqFieldProblemRows)
	reqValidOutput := domain.NewTableDomain(inputTable.GetHeader(), reqValidBody)

	uniqueProblemRows := validatorService.validateUniqueElements(reqValidOutput)
	uniqueValidBody, uniqueInvalidBody := validatorService.filterInvalidIndexesInTableBody(reqValidOutput.GetBody(), uniqueProblemRows)

	invalidBody := append(reqInvalidBody, uniqueInvalidBody...)
	validOutput = domain.NewTableDomain(inputTable.GetHeader(), uniqueValidBody)
	invalidOutput = domain.NewTableDomain(inputTable.GetHeader(), invalidBody)

	return validOutput, invalidOutput
}
