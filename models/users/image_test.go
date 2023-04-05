package users

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

var now = time.Now()
var ims = []UploadedImage{
	{1, "foo", "images/foo1.png", now},
	{2, "foo", "images/foo2.jpeg", now},
	{3, "bar", "images/bar.png", now},
}

func TestImages_ToImageResponse(t *testing.T) {
	res := []ImageResponse{
		{1, "images/foo1.png"},
		{2, "images/foo2.jpeg"},
		{3, "images/bar.png"},
	}
	tests := []struct {
		name string
		im   Images
		want []ImageResponse
	}{
		{"images", Images(ims), res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.im.ToImageResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToImageResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
