package clubs

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

var cont = []ClubContent{
	{1, consts.DummyUUID, "content1"},
	{2, consts.DummyUUID, "content2"},
	{3, consts.DummyUUID, "content3"},
}

func TestContents_ToContentResponse(t *testing.T) {
	res := []ContentResponse{
		{"content1"},
		{"content2"},
		{"content3"},
	}
	tests := []struct {
		name string
		c    Contents
		want []ContentResponse
	}{
		{"content", Contents(cont), res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.ToContentResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToContentResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
