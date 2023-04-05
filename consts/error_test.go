package consts

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestValidateError_Error(t *testing.T) {
	type fields struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"show error of ValidateError", fields{"foo"}, "foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &ValidateError{
				text: tt.fields.text,
			}
			got := err.Error()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Error() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestNewValidateError(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want *ValidateError
	}{
		{"create ValidateError", args{"foo"}, &ValidateError{"foo"}},
	}
	err := NewValidateError("foo")
	opt := cmp.AllowUnexported(*err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := err
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("NewValidateError() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
