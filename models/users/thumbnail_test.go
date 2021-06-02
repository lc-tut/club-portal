package users

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestUploadedThumbnail_ToThumbnailResponse(t *testing.T) {
	type fields struct {
		ThumbnailID uint32
		Path        string
	}
	testFields := []fields{
		{2, "thumbnails/test1.png"},
		{3, "thumbnails/test2.png"},
	}
	res := []ThumbnailResponse{
		{2, "thumbnails/test1.png"},
		{3, "thumbnails/test2.png"},
	}
	tests := []struct {
		name   string
		fields fields
		want   ThumbnailResponse
	}{
		{"user_thumbnail1", testFields[0], res[0]},
		{"user_thumbnail2", testFields[1], res[1]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UploadedThumbnail{
				ThumbnailID: tt.fields.ThumbnailID,
				Path:        tt.fields.Path,
			}
			got := ut.ToThumbnailResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToThumbnailResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
