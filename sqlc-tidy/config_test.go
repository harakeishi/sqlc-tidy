package sqlctidy

import (
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "yml",
			args: args{dir: "./testData"},
			want: "testData/sqlc.yml",
		},
		{
			name:    "no config file",
			args:    args{dir: "./"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConfigPath(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("configPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("configPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
