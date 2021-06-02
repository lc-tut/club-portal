package clubs

import (
	"database/sql"
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

var (
	scheduleRemark = "schedule_remark"
	sch            = []ClubSchedule{
		{1, consts.DummyUUID, 1, "schedule1", sql.NullString{"", false}},
		{2, consts.DummyUUID, 12, "schedule2", sql.NullString{"", false}},
		{3, consts.DummyUUID, 10, "schedule3", sql.NullString{scheduleRemark, true}},
	}
)

func TestSchedules_ToScheduleResponse(t *testing.T) {
	res := []ScheduleResponse{
		{1, "schedule1", nil},
		{12, "schedule2", nil},
		{10, "schedule3", &scheduleRemark},
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
