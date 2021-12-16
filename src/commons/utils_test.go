package commons

import (
	"reflect"
	"testing"
)

func TestTrimSpacesFromElementsInArray(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name    string
		args    args
		wantOut []string
	}{
		{
			name: "Should remove unnecessary spaces at the beginning and at the end of the elements of an array",
			args: args{
				arr: []string{" a ", " b  c", "d   "},
			},
			wantOut: []string{"a", "b  c", "d"},
		},
		{
			name: "Should empty when array passed is empty",
			args: args{
				arr: []string{},
			},
			wantOut: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := TrimSpacesFromElementsInArray(tt.args.arr); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("TrimSpacesFromElementsInArray() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestFindInStringArray(t *testing.T) {
	type args struct {
		arr   []string
		field string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return true if element is found in array",
			args: args{
				arr:   []string{"a", "b", "c"},
				field: "c",
			},
			want: true,
		},
		{
			name: "Should return true if element is found in array after trim spaces",
			args: args{
				arr:   []string{"a", "b", "  c "},
				field: "c",
			},
			want: true,
		},
		{
			name: "Should return false if element is not found in array",
			args: args{
				arr:   []string{"a", "b", "c"},
				field: "d",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindInStringArray(tt.args.arr, tt.args.field); got != tt.want {
				t.Errorf("FindInStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindInIntArray(t *testing.T) {
	type args struct {
		arr   []int
		field int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return true if element is found in array",
			args: args{
				arr:   []int{1, 2, 3},
				field: 3,
			},
			want: true,
		},
		{
			name: "Should return false if element is not found in array",
			args: args{
				arr:   []int{1, 2, 3},
				field: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindInIntArray(tt.args.arr, tt.args.field); got != tt.want {
				t.Errorf("FindInStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindIndexInArray(t *testing.T) {
	type args struct {
		arr   []string
		field string
	}
	tests := []struct {
		name      string
		args      args
		want      bool
		wantIndex int
	}{
		{
			name: "Should return true and element index if element is found in array",
			args: args{
				arr:   []string{"a", "b", "c", "d", "e"},
				field: "e",
			},
			want:      true,
			wantIndex: 4,
		},
		{
			name: "Should return true and element index if element is found in array after trim spaces",
			args: args{
				arr:   []string{"a", "b", "c", "d", "   e "},
				field: "e",
			},
			want:      true,
			wantIndex: 4,
		},
		{
			name: "Should return false and -1 if element is not found in array",
			args: args{
				arr:   []string{"a", "b", "c", "d", "e"},
				field: "f",
			},
			want:      false,
			wantIndex: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotIndex := FindIndexInArray(tt.args.arr, tt.args.field)
			if got != tt.want {
				t.Errorf("FindIndexInArray() got = %v, want %v", got, tt.want)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("FindIndexInArray() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestHasDuplicatedElementsInArray(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return false if doesnt have element duplicated in array",
			args: args{
				arr: []string{"a", "b", "c", "d", "e"},
			},
			want: false,
		},
		{
			name: "Should return true if have element duplicated in array",
			args: args{
				arr: []string{"a", "b", "c", "d", "e", "a"},
			},
			want: true,
		},
		{
			name: "Should return true if have more than one elements duplicated in array",
			args: args{
				arr: []string{"a", "b", "c", "d", "e", "a", "b", "c", "d", "e"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasDuplicatedElementsInArray(tt.args.arr); got != tt.want {
				t.Errorf("HasDuplicatedElementsInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
