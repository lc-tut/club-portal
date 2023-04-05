package clubs

import (
	"database/sql"
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

var (
	rels = Relations([]DetailRelations{
		{consts.DummyUUID, 1, "火曜日", "19:00~21:00", 1, "place1", 1, sql.NullString{String: placeRemark1, Valid: true}, sql.NullString{String: "", Valid: false}},
		{consts.DummyUUID, 2, "金曜日", "19:00~21:00", 2, "place2", 2, sql.NullString{String: "", Valid: false}, sql.NullString{String: timeRemark1, Valid: true}},
		{consts.DummyUUID, 3, "日曜日", "13:00~", 3, "place3", 3, sql.NullString{String: placeRemark2, Valid: true}, sql.NullString{String: timeRemark2, Valid: true}},
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
		{1, consts.DummyUUID, 1, 1, sql.NullString{String: placeRemark1, Valid: true}, sql.NullString{String: "", Valid: false}},
		{2, consts.DummyUUID, 2, 2, sql.NullString{String: "", Valid: false}, sql.NullString{String: timeRemark1, Valid: true}},
		{3, consts.DummyUUID, 3, 3, sql.NullString{String: placeRemark2, Valid: true}, sql.NullString{String: timeRemark2, Valid: true}},
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

func TestRelations_ToActivityDetailResponse(t *testing.T) {
	res := []ActivityDetailResponse{
		{1, "火曜日", "19:00~21:00", nil, 1, "place1", &placeRemark1},
		{2, "金曜日", "19:00~21:00", &timeRemark1, 2, "place2", nil},
		{3, "日曜日", "13:00~", &timeRemark2, 3, "place3", &placeRemark2},
	}
	tests := []struct {
		name string
		r    Relations
		want []ActivityDetailResponse
	}{
		{"relation_response", rels, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.ToActivityDetailResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToActivityDetailResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
