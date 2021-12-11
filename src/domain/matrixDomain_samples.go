package domain

func BuildSampleMatrixDomainData() *MatrixDomain {
	return &MatrixDomain{
		Data: [][]string{
			{"h0", "h1", "h2"},
			{"d00", "d01", "d03"},
			{"d10", "d11", "d13"},
			{"d20", "d21", "d23"},
		},
	}
}
