package domain

import (
	"strings"
)

type (
	TableColumns []*TableColumn

	TableColumn struct {
		Name          string
		Unique        bool
		PossibleWords []string
	}
)

func (col TableColumn) verifyIfAnyPossibleWordMatch(colName string) bool {
	colName = strings.ToLower(colName)
	for _, word := range col.PossibleWords {
		if strings.Contains(colName, word) {
			return true
		}
	}
	return false
}

func (cols TableColumns) GetTableColumnsMatched(colName string) (matches []string) {
	for _, col := range cols {
		if col.verifyIfAnyPossibleWordMatch(colName) {
			matches = append(matches, col.Name)
		}
	}
	return matches
}
