package reader

import (
	"errors"
	"rain-csv-parser/src/domain"
	"rain-csv-parser/src/reader/mocks"
	"reflect"
	"testing"
)

func Test_readerService_Read(t *testing.T) {
	inputPath := "/some/path.anything"
	matrixMocked := domain.BuildSampleMatrixDomainData()

	type fields struct {
		ioStrategy IOStrategy
	}
	type args struct {
		inputPath string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantMatrix *domain.MatrixDomain
		wantErr    bool
	}{
		{
			name: "Should return input matrix when IOStrategy return data correctly",
			fields: fields{
				ioStrategy: func() IOStrategy {
					strategyMocked := &mocks.IOStrategy{}
					strategyMocked.On("Read", inputPath).Return(matrixMocked.Data, nil)
					return strategyMocked
				}(),
			},
			args: args{
				inputPath: inputPath,
			},
			wantMatrix: matrixMocked,
			wantErr:    false,
		},
		{
			name: "Should return error when IOStrategy return some error",
			fields: fields{
				ioStrategy: func() IOStrategy {
					strategyMocked := &mocks.IOStrategy{}
					strategyMocked.On("Read", inputPath).Return(nil, errors.New("read_error"))
					return strategyMocked
				}(),
			},
			args: args{
				inputPath: inputPath,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &readerService{
				ioStrategy: tt.fields.ioStrategy,
			}
			gotMatrix, err := reader.Read(tt.args.inputPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMatrix, tt.wantMatrix) {
				t.Errorf("Read() gotMatrix = %v, want %v", gotMatrix, tt.wantMatrix)
			}
		})
	}
}
