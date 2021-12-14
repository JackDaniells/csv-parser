package main

import (
	"rain-csv-parser/src/constants"
	"rain-csv-parser/src/domain"
	parserdomain "rain-csv-parser/src/parser/domain"
)

const (
	FILE_NAME   = "roster4"
	INPUT_PATH  = "input/" + FILE_NAME
	OUTPUT_PATH = "output/" + FILE_NAME
	FORMAT      = "csv"
)

func createTableColumns() domain.TableColumnSchemas {
	return domain.TableColumnSchemas{
		{
			Name:          constants.HeaderID,
			PossibleWords: []string{"id", "identification", "recognition", "number"},
			Unique:        true,
			Required:      true,
		},
		{
			Name:          constants.HeaderFullName,
			PossibleWords: []string{"n.", "name"},
			Unique:        false,
			Required:      true,
		},
		{
			Name:          constants.HeaderFirstName,
			PossibleWords: []string{"f.", "first"},
			Unique:        false,
			Required:      false,
		},
		{
			Name:          constants.HeaderLastName,
			PossibleWords: []string{"l.", "last"},
			Unique:        false,
			Required:      false,
		},
		{
			Name:          constants.HeaderEmail,
			PossibleWords: []string{"mail", "e-mail", "email"},
			Unique:        true,
			Required:      true,
		},
		{
			Name:          constants.HeaderSalary,
			PossibleWords: []string{"salary", "wage", "pay", "earnings", "income"},
			Unique:        false,
			Required:      true,
		},
	}
}

func createMatcherSelector() parserdomain.MatchSelector {
	return parserdomain.MatchSelector{
		{
			Matches:  []string{constants.HeaderFirstName, constants.HeaderFullName},
			Selected: constants.HeaderFirstName,
		},
		{
			Matches:  []string{constants.HeaderLastName, constants.HeaderFullName},
			Selected: constants.HeaderLastName,
		},
	}
}

func createColumnGrouper() parserdomain.ColumnGrouper {
	return parserdomain.ColumnGrouper{
		{
			Headers:   []string{constants.HeaderFirstName, constants.HeaderLastName},
			GroupName: constants.HeaderFullName,
			Separator: " ",
		},
	}
}
