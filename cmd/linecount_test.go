package cmd

import (
	"io"
	"strings"
	"testing"
)

func Test_checkLineCount(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "null",
			args:    args{reader: strings.NewReader("")},
			want:    0,
			wantErr: false,
		},
		{
			name:    "basic",
			args:    args{reader: strings.NewReader("how do\nyou\nturn this\non\n")},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkLineCount(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkLineCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkLineCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
