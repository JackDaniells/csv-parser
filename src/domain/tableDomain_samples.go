package domain

func BuildSampleMatrixDomainData() *TableDomain {
	return &TableDomain{
		header: []string{"h0", "h1", "h2"},
		body: [][]string{
			{"d00", "d01", "d03"},
			{"d10", "d11", "d13"},
			{"d20", "d21", "d23"},
		},
	}
}
