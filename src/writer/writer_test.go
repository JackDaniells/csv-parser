package writer

import (
	"errors"
	"rain-csv-parser/src/domain"
	"rain-csv-parser/src/writer/mocks"
	"testing"
)

func Test_writerService_Write(t *testing.T) {
	outputPath := "/some/path.anything"
	matrixMocked := domain.BuildSampleMatrixDomainData()

	type fields struct {
		ioStrategy IOStrategy
	}
	type args struct {
		matrix     *domain.MatrixDomain
		outputPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should return input matrix when IOStrategy return data correctly",
			fields: fields{
				ioStrategy: func() IOStrategy {
					strategyMocked := &mocks.IOStrategy{}
					strategyMocked.On("Write", matrixMocked.Data, outputPath).Return(nil)
					return strategyMocked
				}(),
			},
			args: args{
				matrix:     matrixMocked,
				outputPath: outputPath,
			},
			wantErr: false,
		},
		{
			name: "Should return error when IOStrategy return some error",
			fields: fields{
				ioStrategy: func() IOStrategy {
					strategyMocked := &mocks.IOStrategy{}
					strategyMocked.On("Write", matrixMocked.Data, outputPath).Return(errors.New("write_error"))
					return strategyMocked
				}(),
			},
			args: args{
				matrix:     matrixMocked,
				outputPath: outputPath,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &writerService{
				ioStrategy: tt.fields.ioStrategy,
			}
			if err := writer.Write(tt.args.matrix, tt.args.outputPath); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
