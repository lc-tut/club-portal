package clubs

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

func TestClubThumbnail_ToThumbnailResponse(t *testing.T) {
	type fields struct {
		ThumbnailID uint32
		ClubUUID    string
		Path        string
	}
	testField := []fields{
		{1, consts.DummyUUID, consts.DefaultThumbnailPath},
		{2, consts.DummyUUID, "thumbnails/test1.png"},
		{3, consts.DummyUUID, "thumbnails/test2.png"},
	}
	res := []ThumbnailResponse{
		{1, consts.DefaultThumbnailPath},
		{2, "thumbnails/test1.png"},
		{3, "thumbnails/test2.png"},
	}
	tests := []struct {
		name   string
		fields fields
		want   ThumbnailResponse
	}{
		{"thumbnail1", testField[0], res[0]},
		{"thumbnail2", testField[1], res[1]},
		{"thumbnail3", testField[2], res[2]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := &ClubThumbnail{
				ThumbnailID: tt.fields.ThumbnailID,
				ClubUUID:    tt.fields.ClubUUID,
				Path:        tt.fields.Path,
			}
			got := ct.ToThumbnailResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToThumbnailResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
