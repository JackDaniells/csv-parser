package domain

import (
	"reflect"
	"testing"
)

func TestColumnGroup_FindColumnIndexes(t *testing.T) {
	argHeaders := []string{"A", "B", "C", "D", "E"}

	type fields struct {
		Headers   []string
		GroupName string
		Separator string
	}
	type args struct {
		headers []string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantAllFound   bool
		wantColIndexes []int
	}{
		{
			name: "Should return true and array with headers indexes when all headers in group are found",
			fields: fields{
				Headers:   []string{"D", "E"},
				GroupName: "DE",
			},
			args: args{
				headers: argHeaders,
			},
			wantAllFound:   true,
			wantColIndexes: []int{3, 4},
		},
		{
			name: "Should validate if index array returned keep the order of the elements if the first header is positioned after the second in list",
			fields: fields{
				Headers:   []string{"E", "A"},
				GroupName: "DE",
			},
			args: args{
				headers: argHeaders,
			},
			wantAllFound:   true,
			wantColIndexes: []int{4, 0},
		},
		{
			name: "Should return false when some headers in group are not found",
			fields: fields{
				Headers:   []string{"A", "F"},
				GroupName: "AF",
			},
			args: args{
				headers: argHeaders,
			},
			wantAllFound:   false,
			wantColIndexes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := &ColumnGrouper{
				Headers:   tt.fields.Headers,
				GroupName: tt.fields.GroupName,
				Separator: tt.fields.Separator,
			}
			gotAllFound, gotColIndexes := group.FindColumnIndexes(tt.args.headers)
			if gotAllFound != tt.wantAllFound {
				t.Errorf("FindColumnIndexes() gotAllFound = %v, want %v", gotAllFound, tt.wantAllFound)
			}
			if !reflect.DeepEqual(gotColIndexes, tt.wantColIndexes) {
				t.Errorf("FindColumnIndexes() gotColIndexes = %v, want %v", gotColIndexes, tt.wantColIndexes)
			}
		})
	}
}
