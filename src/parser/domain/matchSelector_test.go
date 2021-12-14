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
			name: "Should return true when all matches in the list are in columnMatcher in same order",
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
			name: "Should return true when all matches in the list are in columnMatcher, no matter the element order",
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
			name: "Should return false if some match in the list arent in columnMatcher",
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
			name: "Should return false if any match in the list are in columnMatcher",
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
			colMatcher := columnMatcher{
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
		want     *columnMatcher
	}{
		{
			name: "Should return first element in selector list when all header list for this element is in match list",
			selector: []*columnMatcher{
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
			want: &columnMatcher{
				Matches:  []string{"B", "A"},
				Selected: "B",
			},
		},
		{
			name: "Should return second element in selector list when all header list for this element is in match list",
			selector: []*columnMatcher{
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
			want: &columnMatcher{
				Matches:  []string{"A", "B"},
				Selected: "B",
			},
		},
		{
			name: "Should return nil when any element match with input",
			selector: []*columnMatcher{
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
			selector: []*columnMatcher{
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
			want: &columnMatcher{
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
