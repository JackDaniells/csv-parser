package domain

import (
	"rain-csv-parser/src/commons"
)

type (
	MatchSelector []*columnMatcher

	columnMatcher struct {
		Matches  []string
		Selected string
	}
)

func (colMatcher columnMatcher) allColumnsInMatchesFound(matches []string) bool {
	for _, field := range colMatcher.Matches {
		found := commons.FindInArrayString(matches, field)
		if !found {
			return false
		}
	}
	return true
}

func (selector MatchSelector) GetColumnMatcher(matches []string) *columnMatcher {
	for _, sel := range selector {
		if sel.allColumnsInMatchesFound(matches) {
			return sel
		}
	}
	return nil
}
