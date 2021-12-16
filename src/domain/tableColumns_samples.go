package domain

func BuildSampleTableColumnSchemas() TableColumnSchemas {
	return TableColumnSchemas{
		{
			Name:          "test",
			Unique:        true,
			Required:      false,
			PossibleWords: []string{"test", "tested", "mock", "stub"},
		},
		{
			Name:          "money",
			PossibleWords: []string{"money", "salary", "wage", "pay", "earnings", "income"},
			Unique:        true,
			Required:      true,
		},
		{
			Name:          "money maker",
			PossibleWords: []string{"maker", "functionary", "servant", "worker"},
			Unique:        false,
			Required:      true,
		},
	}
}
