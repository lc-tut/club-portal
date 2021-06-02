package clubs

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var (
	timeRemark1 = "time remark1"
	timeRemark2 = "time remark2"
	times       = []ClubTime{
		{1, "火曜日", "19:00~21:00", nil},
		{2, "金曜日", "19:00~21:00", nil},
		{3, "日曜日", "13:00~", nil},
	}
)

func TestTimes_ToTimeResponse(t *testing.T) {
	type args struct {
		remarks []ClubRemark
	}
	res := []TimeResponse{
		{"火曜日", "19:00~21:00", nil},
		{"金曜日", "19:00~21:00", &timeRemark1},
		{"日曜日", "13:00~", &timeRemark2},
	}
	tests := []struct {
		name string
		t    Times
		args args
		want []TimeResponse
	}{
		{"time_response", Times(times), args{remarks}, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.ToTimeResponse(tt.args.remarks)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToTimeResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
