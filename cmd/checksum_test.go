package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"io"
	"strings"
	"testing"
)

func Test_checksum(t *testing.T) {
	type args struct {
		reader io.Reader
		h      hash.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "md5",
			args: args{
				reader: strings.NewReader("how do\nyou\nturn this\non\n"),
				h:      md5.New(),
			},
			want:    "a8c5d553ed101646036a811772ffbdd8",
			wantErr: false,
		},
		{
			name: "sha1",
			args: args{
				reader: strings.NewReader("how do\nyou\nturn this\non\n"),
				h:      sha1.New(),
			},
			want:    "a656582ca3143a5f48718f4a15e7df018d286521",
			wantErr: false,
		},
		{
			name: "sha256",
			args: args{
				reader: strings.NewReader("how do\nyou\nturn this\non\n"),
				h:      sha256.New(),
			},
			want:    "495a3496cfd90e68a53b5e3ff4f9833b431fe996298f5a28228240ee2a25c09d",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checksum(tt.args.reader, tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("checksum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
