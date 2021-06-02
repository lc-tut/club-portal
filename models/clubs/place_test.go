package clubs

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var (
	placeRemark1 = "place remark1"
	placeRemark2 = "place remark2"
	places       = []ClubPlace{
		{1, "place1", nil},
		{2, "place2", nil},
		{3, "place3", nil},
	}
)

func TestPlaces_ToPlaceResponse(t *testing.T) {
	type args struct {
		remarks []ClubRemark
	}

	res := []PlaceResponse{
		{"place1", &placeRemark1},
		{"place2", nil},
		{"place3", &placeRemark2},
	}
	tests := []struct {
		name string
		p    Places
		args args
		want []PlaceResponse
	}{
		{"place_response", Places(places), args{remarks}, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.ToPlaceResponse(tt.args.remarks)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToPlaceResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
