package clubs

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

var ach = []ClubAchievement{
	{1, consts.DummyUUID, "achievement1"},
	{2, consts.DummyUUID, "achievement2"},
	{3, consts.DummyUUID, "achievement3"},
}

func TestAchievements_ToAchievementResponse(t *testing.T) {
	res := []AchievementResponse{
		{"achievement1"},
		{"achievement2"},
		{"achievement3"},
	}
	tests := []struct {
		name string
		ac   Achievements
		want []AchievementResponse
	}{
		{"achievement", Achievements(ach), res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ac.ToAchievementResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToAchievementResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
