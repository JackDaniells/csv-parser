package csv

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_csvImplementation_CanExecute(t *testing.T) {
	type args struct {
		extension string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return true if extension passed is csv",
			args: args{
				extension: "csv",
			},
			want: true,
		},
		{
			name: "Should return false if extension passed is not csv",
			args: args{
				extension: "xls",
			},
			want: false,
		}, {
			name: "Should return false if extension passed is empty",
			args: args{
				extension: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csvImp := &csvImplementation{}
			if got := csvImp.CanExecute(tt.args.extension); got != tt.want {
				t.Errorf("CanExecute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_csvImplementation_readCSVFile(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Should return matrix when csv data received in buffer parser with success",
			args: args{
				reader: func() io.Reader {
					var buffer bytes.Buffer
					buffer.WriteString("fake,csv,header\nfake,csv,data")
					return &buffer
				}(),
			},
			want: [][]string{
				{"fake", "csv", "header"},
				{"fake", "csv", "data"},
			},
			wantErr: false,
		},
		{
			name: "Should return only one header when csv data received in buffer doesnt have commas",
			args: args{
				reader: func() io.Reader {
					var buffer bytes.Buffer
					buffer.WriteString("Hi Jon-<br> how are you\n")
					return &buffer
				}(),
			},
			want: [][]string{
				{"Hi Jon-<br> how are you"},
			},
			wantErr: false,
		},
		{
			name: "Should return error when data received in buffer is invalid",
			args: args{
				reader: func() io.Reader {
					var buffer bytes.Buffer
					buffer.WriteString("§a,\"b\nc∑\"d,e")
					return &buffer
				}(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csvImp := &csvImplementation{}
			got, err := csvImp.readCSVFile(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCSVFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCSVFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_csvImplementation_writeCSVFile(t *testing.T) {
	type args struct {
		file   io.Writer
		matrix [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should write csv file in buffer successfully",
			args: args{
				file: func() io.Writer {
					var buffer bytes.Buffer
					return &buffer
				}(),
				matrix: [][]string{
					{"fake", "csv", "header"},
					{"fake", "csv", "data"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csvImp := &csvImplementation{}
			file := &bytes.Buffer{}
			err := csvImp.writeCSVFile(file, tt.args.matrix)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeCSVFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_csvImplementation_Read(t *testing.T) {
	type args struct {
		inputPath string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Should read valid input file and return matrix with csv data",
			args: args{
				inputPath: "mocks/valid.csv",
			},
			want: [][]string{
				{"fake", "csv", "header"},
				{"fake", "csv", "body"},
			},
			wantErr: false,
		},
		{
			name: "Should return error when read invalid input file",
			args: args{
				inputPath: "mocks/invalid.csv",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should return error when try open a file that doesn't exist",
			args: args{
				inputPath: "mocks/not_found.csv",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csvImp := &csvImplementation{}
			got, err := csvImp.Read(tt.args.inputPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_csvImplementation_Write(t *testing.T) {
	type args struct {
		matrix     [][]string
		outputPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should write file with success in mock folder",
			args: args{
				matrix: [][]string{
					{"fake", "csv", "header"},
					{"fake", "csv", "body"},
				},
				outputPath: "mocks/out.csv",
			},
			wantErr: false,
		},
		{
			name: "Should return error when try write in a folder that doesn't exist",
			args: args{
				matrix: [][]string{
					{"fake", "csv", "header"},
					{"fake", "csv", "body"},
				},
				outputPath: "not/found/folder/out.csv",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csvImp := &csvImplementation{}
			if err := csvImp.Write(tt.args.matrix, tt.args.outputPath); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
