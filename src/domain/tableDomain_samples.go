package domain

func BuildSampleMatrixDomainData() *TableDomain {
	return &TableDomain{
		Data: [][]string{
			{"h0", "h1", "h2"},
			{"d00", "d01", "d03"},
			{"d10", "d11", "d13"},
			{"d20", "d21", "d23"},
		},
	}
}
