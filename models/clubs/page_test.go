package clubs

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
	"time"
)

var (
	tm             = time.Now()
	campus         = consts.CampusHachioji.ToPrimitive()
	clubType, page = consts.CultureType.ToPrimitive(), Pages([]ClubPage{
		{
			ClubUUID:    consts.DummyUUID,
			ClubSlug:    "aaaaaaaaaaaaaaa",
			Name:        "example",
			Description: "sample desc",
			Campus:      campus,
			ClubType:    clubType,
			Visible:     false,
			UpdatedAt:   tm,
			Thumbnail: ClubThumbnail{
				ThumbnailID: 1,
				ClubUUID:    consts.DummyUUID,
				Path:        consts.DefaultThumbnailPath,
			},
			Contents:     cont,
			Links:        links,
			Schedules:    sch,
			Achievements: ach,
			Videos:       videos,
			ActivityDetails: []ActivityDetail{
				{times[0].TimeID, places[0].PlaceID, consts.DummyUUID, remarks[0]},
				{times[1].TimeID, places[1].PlaceID, consts.DummyUUID, remarks[1]},
				{times[2].TimeID, places[2].PlaceID, consts.DummyUUID, remarks[2]},
			},
		},
	})
)

func TestPages_GetUUIDs(t *testing.T) {
	res := []string{consts.DummyUUID}
	tests := []struct {
		name string
		p    Pages
		want []string
	}{
		{"page_uuids", page, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.GetUUIDs()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetUUIDs() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestPages_ToExternalInfo(t *testing.T) {
	res := []ClubPageExternalInfo{
		{
			ClubUUID:    consts.DummyUUID,
			ClubSlug:    "aaaaaaaaaaaaaaa",
			Name:        "example",
			Description: "sample desc",
			Campus:      campus,
			ClubType:    clubType,
			UpdatedAt:   tm,
			Thumbnail: ThumbnailResponse{
				1, consts.DefaultThumbnailPath,
			},
		},
	}
	tests := []struct {
		name string
		p    Pages
		want []ClubPageExternalInfo
	}{
		{"page_external_info", page, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.ToExternalInfo()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToExternalInfo() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
