package domain

type InputMatrix struct {
	Rows []Row
}

type Row struct {
	Columns []string
}

func (i *InputMatrix) AddRow(row Row) {
	i.Rows = append(i.Rows, row)
}

func (f *InputMatrix) IsEmpty() bool {
	return len(f.Rows) == 0
}

func (f *InputMatrix) Headers() Row {
	return f.Rows[0]
}

func (f *InputMatrix) Body() []Row {
	return f.Rows[1:]
}
