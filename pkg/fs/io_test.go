package fs

import (
	"testing"
)

func TestIsValidFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "non-existent-file",
			args:    args{filepath: "../../tests/non-exist-file.txt"},
			wantErr: true,
		},
		{
			name:    "is-directory",
			args:    args{filepath: "../../tests"},
			wantErr: true,
		},
		{
			name:    "is-file",
			args:    args{filepath: "../../tests/myfile.txt"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsValidFile(tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("IsValidFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsBinary(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "is-binary",
			args:    args{filepath: "../../tests/fops"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "not-binary",
			args:    args{filepath: "../../tests/myfile.txt"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "non-existent-file",
			args:    args{filepath: "../../tests/non-existent-file.txt"},
			want:    false,
			wantErr: true,
		},
		{
			name:    "is-directory",
			args:    args{filepath: "../../tests"},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsBinary(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
