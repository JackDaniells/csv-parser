package parser

import (
	"errors"
	"fmt"
	"rain-csv-parser/src/commons"
	"rain-csv-parser/src/domain"
	parserdomain "rain-csv-parser/src/parser/domain"
	"rain-csv-parser/src/pkg/logger"
	"strings"
)

type parserService struct {
	tableColumns    domain.TableColumnSchemas
	matcherSelector parserdomain.MatchSelector
	columnGrouper   parserdomain.ColumnGrouper
}

func NewParserService(
	tableColumns domain.TableColumnSchemas,
	matcherSelector parserdomain.MatchSelector,
	columnGrouper parserdomain.ColumnGrouper) *parserService {
	return &parserService{
		tableColumns:    tableColumns,
		matcherSelector: matcherSelector,
		columnGrouper:   columnGrouper,
	}
}

func (parserService *parserService) standardizeHeaderName(originalHeaderName string) (string, error) {
	matches := parserService.tableColumns.GetTableColumnsMatched(originalHeaderName)
	matchesLength := len(matches)

	if matchesLength == 0 {
		return originalHeaderName, nil
	}

	if matchesLength == 1 {
		return matches[0], nil
	}

	matcherSelector := parserService.matcherSelector.GetColumnMatcher(matches)
	if matcherSelector == nil {
		errMessage := fmt.Sprintf("ambiguous header found. [name: %s] [matches: %s]",
			originalHeaderName,
			strings.Join(matches, ","))
		return "", errors.New(errMessage)
	}

	return matcherSelector.Selected, nil
}

func (parserService *parserService) standardizeHeaderNames(matrixHeaders []string) (stdHeaders []string, err error) {
	for _, headerName := range matrixHeaders {
		stdHeaderName, err := parserService.standardizeHeaderName(headerName)
		if err != nil {
			return nil, err
		}
		stdHeaders = append(stdHeaders, stdHeaderName)
	}
	return stdHeaders, nil
}

func (parserService *parserService) groupHeaders(headers []string, group *parserdomain.ColumnGroup) (
	newHeader []string) {
	newHeader = append(headers, group.GroupName)
	return newHeader
}

func (parserService *parserService) groupBody(body [][]string, group *parserdomain.ColumnGroup, groupColIndexes []int) (
	newBody [][]string) {
	newBody = body
	for index, row := range body {
		groupRow := ""
		for _, columnIndex := range groupColIndexes {
			groupRow += row[columnIndex] + group.Separator
		}
		groupRow = groupRow[:len(groupRow)-1]
		newBody[index] = append(row, groupRow)
	}
	return newBody
}

func (parserService *parserService) groupTableColumns(headers []string, body [][]string) *domain.TableDomain {
	outputTable := domain.NewTableDomain(headers, body)
	for _, group := range parserService.columnGrouper {
		allIndexesFound, groupIndexes := group.FindColumnIndexes(headers)
		if allIndexesFound {
			outputHeader := parserService.groupHeaders(outputTable.GetHeader(), group)
			outputBody := parserService.groupBody(outputTable.GetBody(), group, groupIndexes)
			outputTable.SetHeader(outputHeader)
			outputTable.SetBody(outputBody)
		}
	}
	return outputTable
}

func (parserService *parserService) validateHeaderDuplication(headers []string) error {
	headers = commons.TrimSpacesFromArray(headers)
	if commons.HasDuplicatedElementsInArray(headers) {
		errMessage := fmt.Sprintf("more than one headers found with same name [headers: %s]",
			strings.Join(headers, ","))
		return errors.New(errMessage)
	}
	return nil
}

func (parserService *parserService) Standardize(inputTable *domain.TableDomain) (outputTable *domain.TableDomain, err error) {
	logger.Debug().Log("Standardizing matrix header names...")
	stdHeaders, err := parserService.standardizeHeaderNames(inputTable.GetHeader())
	if err != nil {
		logger.Error().Log("Error when standardizing header names:", err.Error())
		return nil, err
	}

	logger.Debug().Log("Grouping matrix columns...")
	outputTable = parserService.groupTableColumns(stdHeaders, inputTable.GetBody())

	logger.Debug().Log("Validating if some header is duplicated...")
	err = parserService.validateHeaderDuplication(outputTable.GetHeader())
	if err != nil {
		logger.Error().Log("Error in validation:", err.Error())
		return nil, err
	}

	return outputTable, nil

}
