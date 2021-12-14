package domain

import "errors"

type TableDomain struct {
	Data [][]string
}

func NewMatrixDomain(data [][]string) *TableDomain {
	return &TableDomain{
		Data: data,
	}
}

func (m *TableDomain) IsEmpty() bool {
	return len(m.Data) == 0
}

func (m *TableDomain) GetHeaders() []string {
	return m.Data[0]
}

func (m *TableDomain) SetHeaders(headers []string) {
	m.Data[0] = headers
}

func (m *TableDomain) GetBody() [][]string {
	return m.Data[1:]
}

func (m *TableDomain) HasBody() bool {
	return len(m.GetBody()) != 0
}

func (m *TableDomain) GetRow(index int) ([]string, error) {
	body := m.GetBody()
	if index >= len(body) {
		return nil, errors.New("index out of bounds")
	}
	return body[index], nil
}

func (m *TableDomain) GetColumn(index int) (col []string, err error) {
	body := m.GetBody()
	if index < 0 || index >= len(body[0]) {
		return nil, errors.New("index out of bounds")
	}

	for _, row := range body {
		col = append(col, row[index])
	}
	return col, nil
}
