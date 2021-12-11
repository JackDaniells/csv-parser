package iostrategy

import (
	"testing"
)

func Test_ioStrategySelector_GetStrategy(t *testing.T) {
	type args struct {
		strategy string
	}
	tests := []struct {
		name    string
		args    args
		want    IOStrategy
		wantErr bool
	}{
		{
			name: "Should return IOStrategy implementation when request strategy for CSV files",
			args: args{
				strategy: "csv",
			},
			wantErr: false,
		},
		{
			name: "Should return error when when request strategy not implemented",
			args: args{
				strategy: "pdf",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ioStrategy := &ioStrategySelector{}
			_, err := ioStrategy.GetStrategy(tt.args.strategy)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStrategy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
