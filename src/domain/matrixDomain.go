package domain

type MatrixDomain struct {
	Data [][]string
}

func NewMatrixDomain(data [][]string) *MatrixDomain {
	return &MatrixDomain{
		Data: data,
	}
}

func (m *MatrixDomain) IsEmpty() bool {
	return len(m.Data) == 0
}

func (m *MatrixDomain) Headers() []string {
	return m.Data[0]
}

func (m *MatrixDomain) Body() [][]string {
	return m.Data[1:]
}
