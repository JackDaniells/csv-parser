package domain

import (
	"strings"
)

type (
	TableColumnSchemas []*TableColumnSchema

	TableColumnSchema struct {
		Name          string
		Unique        bool
		Required      bool
		PossibleWords []string
	}
)

func (col TableColumnSchema) verifyIfAnyPossibleWordMatch(colName string) bool {
	colName = strings.ToLower(colName)
	for _, word := range col.PossibleWords {
		if strings.Contains(colName, word) {
			return true
		}
	}
	return false
}

func (cols TableColumnSchemas) GetTableColumnsMatched(colName string) (matches []string) {
	for _, col := range cols {
		if col.verifyIfAnyPossibleWordMatch(colName) {
			matches = append(matches, col.Name)
		}
	}
	return matches
}
