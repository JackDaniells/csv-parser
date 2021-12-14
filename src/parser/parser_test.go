package parser

import (
	"rain-csv-parser/src/domain"
	"reflect"
	"testing"
)

func Test_standardizerService_Standardize(t *testing.T) {
	type args struct {
		inputMatrix *domain.TableDomain
	}
	tests := []struct {
		name              string
		args              args
		wantValidOutput   *domain.TableDomain
		wantInvalidOutput *domain.TableDomain
		wantErr           bool
	}{
		{
			name:              "",
			args:              args{},
			wantValidOutput:   nil,
			wantInvalidOutput: nil,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			standardizerService := &parserService{}
			gotValidOutput, gotInvalidOutput, err := standardizerService.Standardize(tt.args.inputMatrix)
			if (err != nil) != tt.wantErr {
				t.Errorf("Standardize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValidOutput, tt.wantValidOutput) {
				t.Errorf("Standardize() gotValidOutput = %v, want %v", gotValidOutput, tt.wantValidOutput)
			}
			if !reflect.DeepEqual(gotInvalidOutput, tt.wantInvalidOutput) {
				t.Errorf("Standardize() gotInvalidOutput = %v, want %v", gotInvalidOutput, tt.wantInvalidOutput)
			}
		})
	}
}
