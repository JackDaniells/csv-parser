package writer

import (
	"errors"
	"rain-csv-parser/src/domain"
	"rain-csv-parser/src/writer/mocks"
	"testing"
)

func Test_writerService_Write(t *testing.T) {
	outputPath := "/some/path.anything"
	tableMocked := domain.BuildSampleTableDomainData()
	bodyMocked := [][]string{
		{"h0", "h1", "h2"},
		{"d00", "d01", "d03"},
		{"d10", "d11", "d13"},
		{"d20", "d21", "d23"},
	}

	type fields struct {
		ioStrategy IOStrategy
	}
	type args struct {
		matrix     *domain.TableDomain
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
					strategyMocked.On("Write", bodyMocked, outputPath).Return(nil)
					return strategyMocked
				}(),
			},
			args: args{
				matrix:     tableMocked,
				outputPath: outputPath,
			},
			wantErr: false,
		},
		{
			name: "Should return error when IOStrategy return some error",
			fields: fields{
				ioStrategy: func() IOStrategy {
					strategyMocked := &mocks.IOStrategy{}
					strategyMocked.On("Write", bodyMocked, outputPath).Return(errors.New("write_error"))
					return strategyMocked
				}(),
			},
			args: args{
				matrix:     tableMocked,
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
