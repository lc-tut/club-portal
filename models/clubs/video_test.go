package clubs

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

var videos = []ClubVideo{
	{1, consts.DummyUUID, "URL1"},
	{2, consts.DummyUUID, "URL2"},
	{3, consts.DummyUUID, "URL3"},
}

func TestVideos_ToVideoResponse(t *testing.T) {
	res := []VideoResponse{
		{"URL1"},
		{"URL2"},
		{"URL3"},
	}
	tests := []struct {
		name string
		v    Videos
		want []VideoResponse
	}{
		{"video", Videos(videos), res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.ToVideoResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToVideoResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
