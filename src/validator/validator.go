package validator

import (
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
		if !col.Required {
			continue
		}
		colFound, colIndex := commons.FindIndexInArray(headers, col.Name)
		if !colFound || !validatorService.isValidElement(row[colIndex]) {
			return false
		}
	}
	return true
}

func (validatorService *validatorService) validateRequiredFieldsAreFilled(inputTable *domain.TableDomain) ([][]string, [][]string) {
	validBody := [][]string{}
	invalidBody := [][]string{}

	tableHeader := inputTable.GetHeader()
	for _, row := range inputTable.GetBody() {
		if validatorService.checkAllRequiredFieldsAreFilledInRow(tableHeader, row) {
			validBody = append(validBody, row)
		} else {
			invalidBody = append(invalidBody, row)
		}
	}
	return validBody, invalidBody
}

func (validatorService *validatorService) hasDuplicationWithValidBody(validBody [][]string, row []string, colIndex int) bool {
	for _, validRow := range validBody {
		if strings.TrimSpace(validRow[colIndex]) == strings.TrimSpace(row[colIndex]) {
			return true
		}
	}
	return false
}

func (validatorService *validatorService) checkAnyUniqueColumnHaveDuplications(validBody [][]string, headers []string, row []string) bool {
	for _, col := range validatorService.tableColumns {
		if !col.Unique {
			continue
		}

		headerFound, colIndex := commons.FindIndexInArray(headers, col.Name)
		if !headerFound {
			continue
		}

		if validatorService.hasDuplicationWithValidBody(validBody, row, colIndex) {
			return true
		}

	}
	return false
}

func (validatorService *validatorService) validateUniqueElements(inputTable *domain.TableDomain) ([][]string, [][]string) {
	validBody := [][]string{}
	invalidBody := [][]string{}

	for _, row := range inputTable.GetBody() {
		if validatorService.checkAnyUniqueColumnHaveDuplications(validBody, inputTable.GetHeader(), row) {
			invalidBody = append(invalidBody, row)
		} else {
			validBody = append(validBody, row)
		}
	}

	return validBody, invalidBody
}

func (validatorService *validatorService) Validate(inputTable *domain.TableDomain) (
	validOutput *domain.TableDomain, invalidOutput *domain.TableDomain) {

	logger.Info().Log("validating if all required columns are filled...")
	reqValidBody, reqInvalidBody := validatorService.validateRequiredFieldsAreFilled(inputTable)
	reqValidOutput := domain.NewTableDomain(inputTable.GetHeader(), reqValidBody)

	logger.Info().Log("validating if unique columns doesnt have duplications...")
	uniqueValidBody, uniqueInvalidBody := validatorService.validateUniqueElements(reqValidOutput)

	invalidBody := append(reqInvalidBody, uniqueInvalidBody...)
	validOutput = domain.NewTableDomain(inputTable.GetHeader(), uniqueValidBody)
	invalidOutput = domain.NewTableDomain(inputTable.GetHeader(), invalidBody)

	return validOutput, invalidOutput
}
