package domain

import "rain-csv-parser/src/commons"

type (
	GroupAggregator []*ColumnGrouper

	ColumnGrouper struct {
		Headers   []string
		GroupName string
		Separator string
	}
)

func (group *ColumnGrouper) FindColumnIndexes(headers []string) (allFound bool, colIndexes []int) {
	for _, groupColum := range group.Headers {
		found, colIndex := commons.FindIndexInArray(headers, groupColum)
		if !found {
			return false, nil
		}
		colIndexes = append(colIndexes, colIndex)
	}
	return true, colIndexes
}
