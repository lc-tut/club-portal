package clubs

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

var im = []ClubImage{
	{1, consts.DummyUUID, "images/dummy1.png"},
	{2, consts.DummyUUID, "images/dummy2.png"},
	{3, consts.DummyUUID, "images/dummy3.png"},
}

func TestImages_ToImageResponse(t *testing.T) {
	res := []ImageResponse{
		{1, "images/dummy1.png"},
		{2, "images/dummy2.png"},
		{3, "images/dummy3.png"},
	}
	tests := []struct {
		name string
		im   Images
		want []ImageResponse
	}{
		{"image", im, res},
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
