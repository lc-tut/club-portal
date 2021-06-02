package clubs

import (
	"database/sql"
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

var (
	rels = Relations([]DetailRelations{
		{consts.DummyUUID, 1, "火曜日", "19:00~21:00", 1, "place1", 1, sql.NullString{placeRemark1, true}, sql.NullString{"", false}},
		{consts.DummyUUID, 2, "金曜日", "19:00~21:00", 2, "place2", 2, sql.NullString{"", false}, sql.NullString{timeRemark1, true}},
		{consts.DummyUUID, 3, "日曜日", "13:00~", 3, "place3", 3, sql.NullString{placeRemark2, true}, sql.NullString{timeRemark2, true}},
	})
)

func TestRelations_ToClubPlace(t *testing.T) {
	res := []ClubPlace{
		{1, "place1", nil},
		{2, "place2", nil},
		{3, "place3", nil},
	}
	tests := []struct {
		name string
		r    Relations
		want []ClubPlace
	}{
		{"relation_place", rels, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.ToClubPlace()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToClubPlace() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestRelations_ToClubRemark(t *testing.T) {
	res := []ClubRemark{
		{1, consts.DummyUUID, 1, 1, sql.NullString{placeRemark1, true}, sql.NullString{"", false}},
		{2, consts.DummyUUID, 2, 2, sql.NullString{"", false}, sql.NullString{timeRemark1, true}},
		{3, consts.DummyUUID, 3, 3, sql.NullString{placeRemark2, true}, sql.NullString{timeRemark2, true}},
	}
	tests := []struct {
		name string
		r    Relations
		want []ClubRemark
	}{
		{"relation_remark", rels, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.ToClubRemark()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToClubRemark() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestRelations_ToClubTime(t *testing.T) {
	res := []ClubTime{
		{1, "火曜日", "19:00~21:00", nil},
		{2, "金曜日", "19:00~21:00", nil},
		{3, "日曜日", "13:00~", nil},
	}
	tests := []struct {
		name string
		r    Relations
		want []ClubTime
	}{
		{"relation_time", rels, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.ToClubTime()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToClubTime() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
