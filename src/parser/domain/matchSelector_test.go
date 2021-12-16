package domain

import (
	"reflect"
	"testing"
)

func Test_columnMatcher_allColumnsInMatchesFound(t *testing.T) {
	argMatchList := []string{"A", "B"}
	type fields struct {
		Matches  []string
		Selected string
	}
	type args struct {
		matches []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Should return true when all matches in the list are in ColumnMatcher in same order",
			fields: fields{
				Matches:  []string{"A", "B"},
				Selected: "A",
			},
			args: args{
				matches: argMatchList,
			},
			want: true,
		},
		{
			name: "Should return true when all matches in the list are in ColumnMatcher, no matter the element order",
			fields: fields{
				Matches:  []string{"B", "A"},
				Selected: "A",
			},
			args: args{
				matches: argMatchList,
			},
			want: true,
		},
		{
			name: "Should return false if some match in the list arent in ColumnMatcher",
			fields: fields{
				Matches:  []string{"A", "B", "C"},
				Selected: "A",
			},
			args: args{
				matches: argMatchList,
			},
			want: false,
		},
		{
			name: "Should return false if any match in the list are in ColumnMatcher",
			fields: fields{
				Matches:  []string{"C", "D"},
				Selected: "C",
			},
			args: args{
				matches: argMatchList,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			colMatcher := ColumnMatcher{
				Matches:  tt.fields.Matches,
				Selected: tt.fields.Selected,
			}
			if got := colMatcher.allColumnsInMatchesFound(tt.args.matches); got != tt.want {
				t.Errorf("allColumnsInMatchesFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchSelector_GetColumnMatcher(t *testing.T) {
	argMatchList := []string{"A", "B"}
	type args struct {
		matches []string
	}
	tests := []struct {
		name     string
		selector MatchSelector
		args     args
		want     *ColumnMatcher
	}{
		{
			name: "Should return first element in selector list when all header list for this element is in match list",
			selector: []*ColumnMatcher{
				{
					Matches:  []string{"B", "A"},
					Selected: "B",
				},
				{
					Matches:  []string{"B", "C"},
					Selected: "B",
				},
			},
			args: args{
				matches: argMatchList,
			},
			want: &ColumnMatcher{
				Matches:  []string{"B", "A"},
				Selected: "B",
			},
		},
		{
			name: "Should return second element in selector list when all header list for this element is in match list",
			selector: []*ColumnMatcher{
				{
					Matches:  []string{"D", "A"},
					Selected: "B",
				},
				{
					Matches:  []string{"A", "B"},
					Selected: "B",
				},
			},
			args: args{
				matches: argMatchList,
			},
			want: &ColumnMatcher{
				Matches:  []string{"A", "B"},
				Selected: "B",
			},
		},
		{
			name: "Should return nil when any element match with input",
			selector: []*ColumnMatcher{
				{
					Matches:  []string{"C", "A"},
					Selected: "C",
				},
				{
					Matches:  []string{"C", "B"},
					Selected: "B",
				},
			},
			args: args{
				matches: argMatchList,
			},
			want: nil,
		},
		{
			name: "When two elements is in match list can be selected for input argument, should return first one",
			selector: []*ColumnMatcher{
				{
					Matches:  []string{"B", "A"},
					Selected: "B",
				},
				{
					Matches:  []string{"A", "B"},
					Selected: "A",
				},
			},
			args: args{
				matches: argMatchList,
			},
			want: &ColumnMatcher{
				Matches:  []string{"B", "A"},
				Selected: "B",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.selector.GetColumnMatcher(tt.args.matches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetColumnMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
