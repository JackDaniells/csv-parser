package domain

import (
	"rain-csv-parser/src/commons"
)

type (
	MatchSelector []*ColumnMatcher

	ColumnMatcher struct {
		Matches  []string
		Selected string
	}
)

func (colMatcher ColumnMatcher) allColumnsInMatchesFound(matches []string) bool {
	for _, field := range colMatcher.Matches {
		found := commons.FindInStringArray(matches, field)
		if !found {
			return false
		}
	}
	return true
}

func (selector MatchSelector) GetColumnMatcher(matches []string) *ColumnMatcher {
	for _, sel := range selector {
		if sel.allColumnsInMatchesFound(matches) {
			return sel
		}
	}
	return nil
}
