package parser

import (
	"rain-csv-parser/src/domain"
	parserdomain "rain-csv-parser/src/parser/domain"
	"reflect"
	"testing"
)

func createMockedMatchSelector() parserdomain.MatchSelector {
	return parserdomain.MatchSelector{
		{
			Matches:  []string{"money", "money maker"},
			Selected: "money maker",
		},
	}
}

func createMockedColumnGrouper() parserdomain.ColumnGrouper {
	return parserdomain.ColumnGrouper{
		{
			Headers:   []string{"day", "month"},
			GroupName: "date",
			Separator: "/",
		},
	}
}

func Test_parserService_standardizeHeaderName(t *testing.T) {
	type fields struct {
		tableColumns    domain.TableColumnSchemas
		matcherSelector parserdomain.MatchSelector
	}
	type args struct {
		originalHeaderName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should return same header name passed when is not found in any possible words array",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				originalHeaderName: "header named",
			},
			want:    "header named",
			wantErr: false,
		},
		{
			name: "Should convert header name to TableColumnSchema name when is in our possible words array",
			fields: fields{
				tableColumns: domain.BuildSampleTableColumnSchemas(),
			},
			args: args{
				originalHeaderName: "mock",
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "Should search header in matcher selector when input header is found in more than one possible words array",
			fields: fields{
				tableColumns:    domain.BuildSampleTableColumnSchemas(),
				matcherSelector: createMockedMatchSelector(),
			},
			args: args{
				originalHeaderName: "money maker module",
			},
			want:    "money maker",
			wantErr: false,
		},
		{
			name: "Should return error when input header is found in more than one possible words array and doesnt have matcher selector compatible",
			fields: fields{
				tableColumns:    domain.BuildSampleTableColumnSchemas(),
				matcherSelector: createMockedMatchSelector(),
			},
			args: args{
				originalHeaderName: "money mock",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parserService := &parserService{
				tableColumns:    tt.fields.tableColumns,
				matcherSelector: tt.fields.matcherSelector,
			}
			got, err := parserService.standardizeHeaderName(tt.args.originalHeaderName)
			if (err != nil) != tt.wantErr {
				t.Errorf("standardizeHeaderName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("standardizeHeaderName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parserService_standardizeHeaderNames(t *testing.T) {
	type fields struct {
		tableColumns    domain.TableColumnSchemas
		matcherSelector parserdomain.MatchSelector
		columnGrouper   parserdomain.ColumnGrouper
	}
	type args struct {
		matrixHeaders []string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStdHeaders []string
		wantErr        bool
	}{
		{
			name: "Should return same input headers when no one header is found in table column's possible words",
			fields: fields{
				tableColumns:    domain.BuildSampleTableColumnSchemas(),
				matcherSelector: nil,
			},
			args: args{
				matrixHeaders: []string{"bored", "ape", "club"},
			},
			wantStdHeaders: []string{"bored", "ape", "club"},
			wantErr:        false,
		},
		{
			name: "Should convert header names found in table column's possible words",
			fields: fields{
				tableColumns:    domain.BuildSampleTableColumnSchemas(),
				matcherSelector: nil,
			},
			args: args{
				matrixHeaders: []string{"bored", "salary ape", "club"},
			},
			wantStdHeaders: []string{"bored", "money", "club"},
			wantErr:        false,
		},
		{
			name: "Should return error when some header was found in more than one table columns and no have matcherSelector compatible",
			fields: fields{
				tableColumns:    domain.BuildSampleTableColumnSchemas(),
				matcherSelector: nil,
			},
			args: args{
				matrixHeaders: []string{"bored", "salary tested", "club"},
			},
			wantStdHeaders: nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parserService := &parserService{
				tableColumns:    tt.fields.tableColumns,
				matcherSelector: tt.fields.matcherSelector,
				columnGrouper:   tt.fields.columnGrouper,
			}
			gotStdHeaders, err := parserService.standardizeHeaderNames(tt.args.matrixHeaders)
			if (err != nil) != tt.wantErr {
				t.Errorf("standardizeHeaderNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStdHeaders, tt.wantStdHeaders) {
				t.Errorf("standardizeHeaderNames() gotStdHeaders = %v, want %v", gotStdHeaders, tt.wantStdHeaders)
			}
		})
	}
}

func Test_parserService_groupTableColumns(t *testing.T) {
	argBody := [][]string{
		{"d00", "d01", "d03"},
		{"d10", "d11", "d13"},
		{"d20", "d21", "d23"},
	}
	argBodyGrouped := [][]string{
		{"d00", "d01", "d03", "d00/d01"},
		{"d10", "d11", "d13", "d10/d11"},
		{"d20", "d21", "d23", "d20/d21"},
	}

	argHeaderNoGrouper := []string{"some", "simple", "test"}

	type fields struct {
		columnGrouper parserdomain.ColumnGrouper
	}
	type args struct {
		headers []string
		body    [][]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *domain.TableDomain
	}{
		{
			name: "Should return table in same structure when no have grouper for any header",
			fields: fields{
				columnGrouper: createMockedColumnGrouper(),
			},
			args: args{
				headers: argHeaderNoGrouper,
				body:    argBody,
			},
			want: domain.NewTableDomain(argHeaderNoGrouper, argBody),
		},
		{
			name: "Should group header and body when have grouper for two or more headers",
			fields: fields{
				columnGrouper: createMockedColumnGrouper(),
			},
			args: args{
				headers: []string{"day", "month", "test"},
				body:    argBody,
			},
			want: domain.NewTableDomain([]string{"day", "month", "test", "date"}, argBodyGrouped),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parserService := &parserService{
				columnGrouper: tt.fields.columnGrouper,
			}
			if got := parserService.groupTableColumns(tt.args.headers, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupTableColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parserService_validateHeaderDuplication(t *testing.T) {
	type args struct {
		headers []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should return ok when doesnt have headers duplicated",
			args: args{
				headers: []string{"some", "simple", "header", "test"},
			},
			wantErr: false,
		},
		{
			name: "Should return error when some header is duplicated",
			args: args{
				headers: []string{"some", "simple", "header", "test", "some"},
			},
			wantErr: true,
		},
		{
			name: "Should return error when some header is duplicated before trim spaces",
			args: args{
				headers: []string{"some", "simple", "header", "test", "some "},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parserService := &parserService{}
			if err := parserService.validateHeaderDuplication(tt.args.headers); (err != nil) != tt.wantErr {
				t.Errorf("validateHeaderDuplication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parserService_Standardize(t *testing.T) {
	correctHeader := []string{"some", "simple", "header"}
	headerDuplicated := []string{"test", "mock", "header"}
	headerErrorMatch := []string{"test money", "mock", "header"}
	argBody := [][]string{
		{"d00", "d01", "d03"},
		{"d10", "d11", "d13"},
		{"d20", "d21", "d23"},
	}
	type fields struct {
		tableColumns    domain.TableColumnSchemas
		matcherSelector parserdomain.MatchSelector
		columnGrouper   parserdomain.ColumnGrouper
	}
	type args struct {
		inputTable *domain.TableDomain
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantOutputTable *domain.TableDomain
		wantErr         bool
	}{
		{
			name: "Should return output table when standardize runs with no errors",
			fields: fields{
				tableColumns:    domain.BuildSampleTableColumnSchemas(),
				matcherSelector: createMockedMatchSelector(),
				columnGrouper:   createMockedColumnGrouper(),
			},
			args: args{
				inputTable: domain.NewTableDomain(correctHeader, argBody),
			},
			wantOutputTable: domain.NewTableDomain(correctHeader, argBody),
			wantErr:         false,
		},
		{
			name: "Should return error when some exception occurs in match selector",
			fields: fields{
				tableColumns:    domain.BuildSampleTableColumnSchemas(),
				matcherSelector: createMockedMatchSelector(),
				columnGrouper:   createMockedColumnGrouper(),
			},
			args: args{
				inputTable: domain.NewTableDomain(headerErrorMatch, argBody),
			},
			wantOutputTable: nil,
			wantErr:         true,
		},
		{
			name: "Should return error when some exception occurs in duplication validation",
			fields: fields{
				tableColumns:    domain.BuildSampleTableColumnSchemas(),
				matcherSelector: createMockedMatchSelector(),
				columnGrouper:   createMockedColumnGrouper(),
			},
			args: args{
				inputTable: domain.NewTableDomain(headerDuplicated, argBody),
			},
			wantOutputTable: nil,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parserService := &parserService{
				tableColumns:    tt.fields.tableColumns,
				matcherSelector: tt.fields.matcherSelector,
				columnGrouper:   tt.fields.columnGrouper,
			}
			gotOutputTable, err := parserService.Standardize(tt.args.inputTable)
			if (err != nil) != tt.wantErr {
				t.Errorf("Standardize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutputTable, tt.wantOutputTable) {
				t.Errorf("Standardize() gotOutputTable = %v, want %v", gotOutputTable, tt.wantOutputTable)
			}
		})
	}
}
