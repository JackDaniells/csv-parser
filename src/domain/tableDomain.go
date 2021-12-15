package domain

import "errors"

type TableDomain struct {
	header []string
	body   [][]string
}

func NewTableDomain(header []string, body [][]string) *TableDomain {
	return &TableDomain{
		header: header,
		body:   body,
	}
}

func (m *TableDomain) GetHeader() []string {
	return m.header
}

func (m *TableDomain) SetHeader(headers []string) {
	m.header = headers
}

func (m *TableDomain) GetBody() [][]string {
	return m.body
}

func (m *TableDomain) SetBody(body [][]string) {
	m.body = body
}

func (m *TableDomain) GetRow(index int) ([]string, error) {
	if index >= len(m.body) {
		return nil, errors.New("index out of bounds")
	}
	return m.body[index], nil
}

func (m *TableDomain) GetColumn(index int) (col []string, err error) {
	if index < 0 || index >= len(m.body[0]) {
		return nil, errors.New("index out of bounds")
	}

	for _, row := range m.body {
		col = append(col, row[index])
	}
	return col, nil
}

func (m *TableDomain) GetStringMatrixOutput() [][]string {
	outputMatrix := [][]string{
		m.GetHeader(),
	}
	outputMatrix = append(outputMatrix, m.GetBody()...)
	return outputMatrix
}
