package validator

import (
	"errors"
	"fmt"
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

func (validatorService *validatorService) checkDuplicatedFieldsForUniqueColumn(table *domain.TableDomain, column string) (
	duplicatedElements [][]int, err error) {
	columnFound, idColIndex := commons.FindIndexInArray(table.GetHeader(), column)
	if !columnFound {
		return nil, errors.New("unique column not found")
	}
	colData, err := table.GetColumn(idColIndex)
	if err != nil {
		return nil, err
	}
	colData = commons.TrimSpacesFromArray(colData)
	duplicatedElements = commons.GetDuplicatedElementsIndexesInArray(colData)
	return duplicatedElements, nil
}

func (validatorService *validatorService) invalidateDuplicatedRowIndexes(duplicateds [][]int) []int {
	invalidIndexes := []int{}
	for _, row := range duplicateds {
		for i, cell := range row {
			// maintain first element and discard others
			if i > 0 {
				invalidIndexes = append(invalidIndexes, cell)
			}
		}
	}
	return invalidIndexes
}

func (validatorService *validatorService) validateUniqueElements(table *domain.TableDomain) (invalidIndexes []int, err error) {
	for _, col := range validatorService.tableColumns {
		if col.Unique {
			duplicateds, err := validatorService.checkDuplicatedFieldsForUniqueColumn(table, col.Name)
			if err != nil {
				return nil, err
			}
			invalidIndexes = append(invalidIndexes, validatorService.invalidateDuplicatedRowIndexes(duplicateds)...)
		}
	}
	return commons.RemoveDuplicatedFields(invalidIndexes), nil
}

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
	reqFieldProblemRows := validatorService.validateRequiredFieldsAreFilled(inputTable)
	reqValidBody, reqInvalidBody := validatorService.filterTableBody(inputTable.GetBody(), reqFieldProblemRows)

	reqValidOutput := domain.NewTableDomain(inputTable.GetHeader(), reqValidBody)

	uniqueProblemRows, err := validatorService.validateUniqueElements(reqValidOutput)
	if err != nil {
		logger.Info().Log(fmt.Sprintf("error when validate unique elements: %s", err.Error()))
		return nil, nil, err
	}

	uniqueValidBody, uniqueInvalidBody := validatorService.filterTableBody(reqValidOutput.GetBody(), uniqueProblemRows)
	invalidBody := append(reqInvalidBody, uniqueInvalidBody...)
	validOutput = domain.NewTableDomain(inputTable.GetHeader(), uniqueValidBody)
	invalidOutput = domain.NewTableDomain(inputTable.GetHeader(), invalidBody)

	return validOutput, invalidOutput, nil

}
