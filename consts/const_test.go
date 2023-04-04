package consts

import "testing"

func TestCampusType_ToPrimitive(t *testing.T) {
	tests := []struct {
		name string
		ct   CampusType
		want uint8
	}{
		{"kamata", CampusHachioji, 0},
		{"hachioji", CampusKamata, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ct.ToPrimitive(); got != tt.want {
				t.Errorf("ToPrimitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubType_ToPrimitive(t *testing.T) {
	tests := []struct {
		name string
		ct   ClubType
		want uint8
	}{
		{"sports", SportsType, 0},
		{"culture", CultureType, 1},
		{"Kokasai", KokasaiType, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ct.ToPrimitive(); got != tt.want {
				t.Errorf("ToPrimitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserType_ToPrimitive(t *testing.T) {
	tests := []struct {
		name string
		ut   UserType
		want string
	}{
		{"domain", DomainUser, "domain"},
		{"general", GeneralUser, "general"},
		{"admin", AdminUser, "admin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ut.ToPrimitive(); got != tt.want {
				t.Errorf("ToPrimitive() = %v, want %v", got, tt.want)
			}
		})
	}
}
