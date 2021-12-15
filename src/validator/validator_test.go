package validator

import (
	"rain-csv-parser/src/domain"
	"reflect"
	"testing"
)

func Test_validatorService_checkAllRequiredFieldsAreFilledInRow(t *testing.T) {
	type fields struct {
		tableColumns domain.TableColumnSchemas
	}
	type args struct {
		headers []string
		row     []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Should return true when all required fields are filled",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				headers: []string{"money", "money maker", "test"},
				row:     []string{"d00", "d01", "d03"},
			},
			want: true,
		},
		{
			name: "Should return false when some required field are not filled",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				headers: []string{"money", "money maker", "test"},
				row:     []string{"", "d01", "d03"},
			},
			want: false,
		},
		{
			name: "Should return false when some required field are not filled before trim validation",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				headers: []string{"money", "money maker", "test"},
				row:     []string{"   ", "d01", "d03"},
			},
			want: false,
		},
		{
			name: "Should return false when some required header not exist",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				headers: []string{"header", "money maker", "test"},
				row:     []string{"d00", "d01", "d03"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatorService := &validatorService{
				tableColumns: tt.fields.tableColumns,
			}
			if got := validatorService.checkAllRequiredFieldsAreFilledInRow(tt.args.headers, tt.args.row); got != tt.want {
				t.Errorf("checkAllRequiredFieldsAreFilledInRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatorService_validateRequiredFieldsAreFilled(t *testing.T) {
	type fields struct {
		tableColumns domain.TableColumnSchemas
	}
	type args struct {
		inputTable *domain.TableDomain
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantInvalidRows []int
	}{
		{
			name: "Should return empty array when all cells required are filled",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				inputTable: func() *domain.TableDomain {
					mockHeader := []string{"money", "money maker", "test"}
					mockBody := domain.BuildSampleTableBody()
					return domain.NewTableDomain(mockHeader, mockBody)
				}(),
			},
			wantInvalidRows: nil,
		},
		{
			name: "Should return array with all indexes of rows with required are not filled",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				inputTable: func() *domain.TableDomain {
					mockHeader := []string{"money", "money maker", "test"}
					mockBody := [][]string{
						{"", "d01", "d03"},
						{"d10", "d11", "d13"},
						{"", "d21", "d23"},
					}
					return domain.NewTableDomain(mockHeader, mockBody)
				}(),
			},
			wantInvalidRows: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatorService := &validatorService{
				tableColumns: tt.fields.tableColumns,
			}
			if gotInvalidRows := validatorService.validateRequiredFieldsAreFilled(tt.args.inputTable); !reflect.DeepEqual(gotInvalidRows, tt.wantInvalidRows) {
				t.Errorf("validateRequiredFieldsAreFilled() = %v, want %v", gotInvalidRows, tt.wantInvalidRows)
			}
		})
	}
}

func Test_validatorService_checkDuplicatedFieldsForUniqueColumn(t *testing.T) {
	type args struct {
		table      *domain.TableDomain
		columnName string
	}
	tests := []struct {
		name                   string
		args                   args
		wantDuplicatedElements [][]int
		wantErr                bool
	}{
		{
			name: "Should return empty duplicated elements array when doesnt have duplications in column",
			args: args{
				table: func() *domain.TableDomain {
					mockHeader := []string{"money", "money maker", "test"}
					mockBody := domain.BuildSampleTableBody()
					return domain.NewTableDomain(mockHeader, mockBody)
				}(),
				columnName: "money",
			},
			wantDuplicatedElements: [][]int{},
			wantErr:                false,
		},
		{
			name: "Should return bidireccional array with indexes of duplicated elements in array",
			args: args{
				table: func() *domain.TableDomain {
					mockHeader := []string{"money", "money maker", "test"}
					mockBody := [][]string{
						{"d00", "d01", "d03"},
						{"d10", "d11", "d13"},
						{"d00", "d21", "d23"},
						{"d10", "d31", "d33"},
					}
					return domain.NewTableDomain(mockHeader, mockBody)
				}(),
				columnName: "money",
			},
			wantDuplicatedElements: [][]int{
				{0, 2},
				{1, 3},
			},
			wantErr: false,
		},
		{
			name: "Should return error when header not found in table",
			args: args{
				table: func() *domain.TableDomain {
					mockHeader := []string{"money", "money maker", "test"}
					mockBody := domain.BuildSampleTableBody()
					return domain.NewTableDomain(mockHeader, mockBody)
				}(),
				columnName: "header",
			},
			wantDuplicatedElements: nil,
			wantErr:                true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatorService := &validatorService{}
			gotDuplicatedElements, err := validatorService.checkDuplicatedFieldsForUniqueColumn(tt.args.table, tt.args.columnName)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkDuplicatedFieldsForUniqueColumn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDuplicatedElements, tt.wantDuplicatedElements) {
				t.Errorf("checkDuplicatedFieldsForUniqueColumn() gotDuplicatedElements = %v, want %v", gotDuplicatedElements, tt.wantDuplicatedElements)
			}
		})
	}
}
