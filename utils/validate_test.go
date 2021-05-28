package utils

import (
	"database/sql"
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

func TestNullStringToStringP(t *testing.T) {
	str := "foo"
	type args struct {
		s sql.NullString
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{"valid value", args{sql.NullString{String: str, Valid: true}}, &str},
		{"invalid value", args{sql.NullString{String: "", Valid: false}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NullStringToStringP(tt.args.s)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NullStringToStringP() = %v, want %v,\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestStringPToString(t *testing.T) {
	str := "foo"
	type args struct {
		s *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"non-nil value", args{&str}, str},
		{"nil value", args{nil}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringPToString(tt.args.s)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("StringPToString() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestStringToNullString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{"valid value", args{"foo"}, sql.NullString{String: "foo", Valid: true}},
		{"invalid value", args{""}, sql.NullString{String: "", Valid: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringToNullString(tt.args.s)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("StringToNullString() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestToCampusType(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name    string
		args    args
		want    consts.CampusType
		wantErr bool
	}{
		{"kamata", args{0}, consts.CampusKamata, false},
		{"hachioji", args{1}, consts.CampusHachioji, false},
		{"invalid value", args{42}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToCampusType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCampusType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToCampusType() got = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestToClubType(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name    string
		args    args
		want    consts.ClubType
		wantErr bool
	}{
		{"sports", args{0}, consts.SportsType, false},
		{"culture", args{1}, consts.CultureType, false},
		{"kokasai", args{2}, consts.KokasaiType, false},
		{"invalid value", args{42}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToClubType(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToClubType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToClubType() got = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestToUserType(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    consts.UserType
		wantErr bool
	}{
		{"admin", args{"admin"}, consts.AdminUser, false},
		{"general", args{"general"}, consts.GeneralUser, false},
		{"domain", args{"domain"}, consts.DomainUser, false},
		{"invalid value", args{"invalid"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUserType(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUserType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToUserType() got = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
