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
			name: "Should return false when some required field are not filled before trim spaces",
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
		wantValidBody   [][]string
		wantInvalidBody [][]string
	}{
		{
			name: "Should return all rows valid when all cells required are filled",
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
			wantValidBody:   domain.BuildSampleTableBody(),
			wantInvalidBody: [][]string{},
		},
		{
			name: "Should return array rows in invalid body matrix when required are not filled",
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
			wantValidBody: [][]string{
				{"d10", "d11", "d13"},
			},
			wantInvalidBody: [][]string{
				{"", "d01", "d03"},
				{"", "d21", "d23"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatorService := &validatorService{
				tableColumns: tt.fields.tableColumns,
			}
			validBody, invalidBody := validatorService.validateRequiredFieldsAreFilled(tt.args.inputTable)
			if !reflect.DeepEqual(validBody, tt.wantValidBody) {
				t.Errorf("validateRequiredFieldsAreFilled() got = %v, want %v", validBody, tt.wantValidBody)
			}
			if !reflect.DeepEqual(invalidBody, tt.wantInvalidBody) {
				t.Errorf("validateRequiredFieldsAreFilled() got1 = %v, want %v", invalidBody, tt.wantInvalidBody)
			}
		})
	}
}

func Test_validatorService_checkAnyUniqueColumnHaveDuplications(t *testing.T) {
	type fields struct {
		tableColumns domain.TableColumnSchemas
	}
	type args struct {
		validatedBody [][]string
		headers       []string
		row           []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Should return false when doesnt have unique fields duplication",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				headers: []string{"money", "money maker", "test"},
				row:     []string{"d00", "d01", "d03"},
				validatedBody: [][]string{
					{"d01", "d11", "d13"},
					{"d02", "d11", "d23"},
					{"d03", "d21", "d33"},
				},
			},
			want: false,
		},
		{
			name: "Should return true when some have unique field has duplication",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				headers: []string{"money", "money maker", "test"},
				row:     []string{"d01", "d01", "d03"},
				validatedBody: [][]string{
					{"d01", "d11", "d13"},
					{"d02", "d11", "d23"},
					{"d03", "d21", "d33"},
				},
			},
			want: true,
		},
		{
			name: "Should return true when some have unique field has duplication before trim spaces",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				headers: []string{"money", "money maker", "test"},
				row:     []string{"d01 ", "d01", "d03"},
				validatedBody: [][]string{
					{"d01", "d11", "d13"},
					{"d02", "d11", "d23"},
					{"d03", "d21", "d33"},
				},
			},
			want: true,
		},
		{
			name: "Should return false when unique header not exist",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				headers: []string{"header", "money maker", "all"},
				row:     []string{"d01", "d11", "d13"},
				validatedBody: [][]string{
					{"d01", "d11", "d13"},
					{"d02", "d11", "d23"},
					{"d03", "d21", "d33"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatorService := &validatorService{
				tableColumns: tt.fields.tableColumns,
			}
			if got := validatorService.checkAnyUniqueColumnHaveDuplications(tt.args.validatedBody, tt.args.headers, tt.args.row); got != tt.want {
				t.Errorf("checkAnyUniqueColumnHaveDuplications() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatorService_validateUniqueElements(t *testing.T) {
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
		wantValidBody   [][]string
		wantInvalidBody [][]string
	}{
		{
			name: "Should return all rows valid when all unique cells doesnt have duplications",
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
			wantValidBody:   domain.BuildSampleTableBody(),
			wantInvalidBody: [][]string{},
		},
		{
			name: "Should return array rows in invalid body matrix when unique cells have duplications",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				inputTable: func() *domain.TableDomain {
					mockHeader := []string{"money", "money maker", "test"}
					mockBody := [][]string{
						{"d00", "d01", "d03"},
						{"d10", "d11", "d13"},
						{"d20", "d21", "d03"},
					}
					return domain.NewTableDomain(mockHeader, mockBody)
				}(),
			},
			wantValidBody: [][]string{
				{"d00", "d01", "d03"},
				{"d10", "d11", "d13"},
			},
			wantInvalidBody: [][]string{
				{"d20", "d21", "d03"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatorService := &validatorService{
				tableColumns: tt.fields.tableColumns,
			}
			got, got1 := validatorService.validateUniqueElements(tt.args.inputTable)
			if !reflect.DeepEqual(got, tt.wantValidBody) {
				t.Errorf("validateUniqueElements() got = %v, want %v", got, tt.wantValidBody)
			}
			if !reflect.DeepEqual(got1, tt.wantInvalidBody) {
				t.Errorf("validateUniqueElements() got1 = %v, want %v", got1, tt.wantInvalidBody)
			}
		})
	}
}

func Test_validatorService_Validate(t *testing.T) {
	sampleInputTable := domain.BuildSampleTableDomainData()
	type fields struct {
		tableColumns domain.TableColumnSchemas
	}
	type args struct {
		inputTable *domain.TableDomain
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantValidOutput   *domain.TableDomain
		wantInvalidOutput *domain.TableDomain
		wantErr           bool
	}{
		{
			name: "Should return all input table as valid when no have table columns required or unique",
			fields: fields{
				tableColumns: domain.TableColumnSchemas{},
			},
			args: args{
				inputTable: sampleInputTable,
			},
			wantValidOutput:   sampleInputTable,
			wantInvalidOutput: domain.NewTableDomain(sampleInputTable.GetHeader(), [][]string{}),
			wantErr:           false,
		},
		{
			name: "Should return all input table as invalid when some table column required are not filled",
			fields: fields{
				tableColumns: domain.TableColumnSchemas{
					{
						Name:     "test",
						Unique:   false,
						Required: true,
					},
				},
			},
			args: args{
				inputTable: sampleInputTable,
			},
			wantValidOutput:   domain.NewTableDomain(sampleInputTable.GetHeader(), [][]string{}),
			wantInvalidOutput: sampleInputTable,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatorService := &validatorService{
				tableColumns: tt.fields.tableColumns,
			}
			gotValidOutput, gotInvalidOutput := validatorService.Validate(tt.args.inputTable)
			if !reflect.DeepEqual(gotValidOutput, tt.wantValidOutput) {
				t.Errorf("Validate() gotValidOutput = %v, want %v", gotValidOutput, tt.wantValidOutput)
			}
			if !reflect.DeepEqual(gotInvalidOutput, tt.wantInvalidOutput) {
				t.Errorf("Validate() gotInvalidOutput = %v, want %v", gotInvalidOutput, tt.wantInvalidOutput)
			}
		})
	}
}
