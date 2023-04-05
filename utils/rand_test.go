package utils

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"path/filepath"
	"testing"
)

func TestGenerateCSRFState(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"generate", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateCSRFState()
			if err != nil {
				t.Errorf("GenerateCSRFState() occurred error = %v", err)
				return
			}
			length := len(got)
			if length != 32 {
				t.Errorf("GenerateCSRFState() should be 16 bytes (32 lengths); length = %v", length)
			}
		})
	}
}

func TestGenerateFileName(t *testing.T) {
	type args struct {
		fn string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"dummy file", args{"dummy.png"}, "94082d1adc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateFileName(tt.args.fn)
			if err != nil {
				t.Errorf("GenerateFileName() error = %v", err)
				return
			}
			ext := filepath.Ext(tt.args.fn)
			if got[:10] != tt.want && got[18:] != ext {
				t.Errorf("GenerateFileName() got = %v, want %v, ext = %v", got, tt.want, ext)
			}
		})
	}
}

func TestGenerateSlug(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"dummy uuid", args{consts.DummyUUID}, "43677a3769d2536"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateSlug(tt.args.uuid)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GenerateSlug() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
