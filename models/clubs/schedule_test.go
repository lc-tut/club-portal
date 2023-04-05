package clubs

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

var (
	sch = []ClubSchedule{
		{1, consts.DummyUUID, 1, "schedule1"},
		{2, consts.DummyUUID, 12, "schedule2"},
		{3, consts.DummyUUID, 10, "schedule3"},
	}
)

func TestSchedules_ToScheduleResponse(t *testing.T) {
	res := []ScheduleResponse{
		{1, "schedule1"},
		{12, "schedule2"},
		{10, "schedule3"},
	}
	tests := []struct {
		name string
		s    Schedules
		want []ScheduleResponse
	}{
		{"schedule", Schedules(sch), res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.ToScheduleResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToScheduleResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
