package users

import (
	"database/sql"
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

var (
	adminUser   = AdminUser{consts.DummyUUID, "admin@example.com", "admin example"}
	generalUser = GeneralUser{consts.DummyUUID, "general@example.com", "general example", sql.NullString{}}
	domainUser  = DomainUser{consts.DummyUUID, "domain@example.com", "domain example", nil}
)

func TestAdminUser_GetEmail(t *testing.T) {
	type fields struct {
		UserUUID string
		Email    string
		Name     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"admin_email", fields(adminUser), "admin@example.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AdminUser{
				UserUUID: tt.fields.UserUUID,
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
			}
			got := u.GetEmail()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetEmail() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestAdminUser_GetName(t *testing.T) {
	type fields struct {
		UserUUID string
		Email    string
		Name     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"admin_name", fields(adminUser), "admin example"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AdminUser{
				UserUUID: tt.fields.UserUUID,
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
			}
			got := u.GetName()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetName() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestAdminUser_GetRole(t *testing.T) {
	type fields struct {
		UserUUID string
		Email    string
		Name     string
	}
	tests := []struct {
		name   string
		fields fields
		want   consts.UserType
	}{
		{"admin_role", fields(adminUser), "admin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AdminUser{
				UserUUID: tt.fields.UserUUID,
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
			}
			got := u.GetRole()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetRole() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestAdminUser_GetUserID(t *testing.T) {
	type fields struct {
		UserUUID string
		Email    string
		Name     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"admin_uuid", fields(adminUser), consts.DummyUUID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AdminUser{
				UserUUID: tt.fields.UserUUID,
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
			}
			got := u.GetUserID()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetUserID() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestAdminUser_ToUserResponse(t *testing.T) {
	type fields struct {
		UserUUID string
		Email    string
		Name     string
	}
	tests := []struct {
		name   string
		fields fields
		want   *UserResponse
	}{
		{"admin_response", fields(adminUser), &UserResponse{consts.DummyUUID, "admin@example.com", "admin example", "admin"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AdminUser{
				UserUUID: tt.fields.UserUUID,
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
			}
			got := u.ToUserResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToUserResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestDomainUser_GetEmail(t *testing.T) {
	type fields struct {
		UserUUID  string
		Email     string
		Name      string
		Favorites []FavoriteClub
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"domain_email", fields(domainUser), "domain@example.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DomainUser{
				UserUUID:  tt.fields.UserUUID,
				Email:     tt.fields.Email,
				Name:      tt.fields.Name,
				Favorites: tt.fields.Favorites,
			}
			got := u.GetEmail()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetEmail() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestDomainUser_GetName(t *testing.T) {
	type fields struct {
		UserUUID  string
		Email     string
		Name      string
		Favorites []FavoriteClub
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"domain_name", fields(domainUser), "domain example"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DomainUser{
				UserUUID:  tt.fields.UserUUID,
				Email:     tt.fields.Email,
				Name:      tt.fields.Name,
				Favorites: tt.fields.Favorites,
			}
			got := u.GetName()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetName() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestDomainUser_GetRole(t *testing.T) {
	type fields struct {
		UserUUID  string
		Email     string
		Name      string
		Favorites []FavoriteClub
	}
	tests := []struct {
		name   string
		fields fields
		want   consts.UserType
	}{
		{"domain_role", fields(domainUser), "domain"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DomainUser{
				UserUUID:  tt.fields.UserUUID,
				Email:     tt.fields.Email,
				Name:      tt.fields.Name,
				Favorites: tt.fields.Favorites,
			}
			got := u.GetRole()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetRole() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestDomainUser_GetUserID(t *testing.T) {
	type fields struct {
		UserUUID  string
		Email     string
		Name      string
		Favorites []FavoriteClub
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"domain_uuid", fields(domainUser), consts.DummyUUID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DomainUser{
				UserUUID:  tt.fields.UserUUID,
				Email:     tt.fields.Email,
				Name:      tt.fields.Name,
				Favorites: tt.fields.Favorites,
			}
			got := u.GetUserID()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetUserID() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestDomainUser_ToUserResponse(t *testing.T) {
	type fields struct {
		UserUUID  string
		Email     string
		Name      string
		Favorites []FavoriteClub
	}
	tests := []struct {
		name   string
		fields fields
		want   *UserResponse
	}{
		{"domain_response", fields(domainUser), &UserResponse{consts.DummyUUID, "domain@example.com", "domain example", "domain"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DomainUser{
				UserUUID:  tt.fields.UserUUID,
				Email:     tt.fields.Email,
				Name:      tt.fields.Name,
				Favorites: tt.fields.Favorites,
			}
			got := u.ToUserResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToUserResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestGeneralUserSlice_GetEmails(t *testing.T) {
	userSlice := GeneralUserSlice([]GeneralUser{
		generalUser,
		{"bbbbbbbb-bbbb-4bbb-bbbb-bbbbbbbbbbbb", "general2@example.com", "general2", sql.NullString{}},
		{"cccccccc-cccc-4ccc-cccc-cccccccccccc", "general3@example.com", "general3", sql.NullString{}},
	})
	tests := []struct {
		name string
		g    GeneralUserSlice
		want []string
	}{
		{"general_slice_emails", userSlice, []string{"general@example.com", "general2@example.com", "general3@example.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.g.GetEmails()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetEmails() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestGeneralUser_GetEmail(t *testing.T) {
	type fields struct {
		UserUUID string
		Email    string
		Name     string
		ClubUUID sql.NullString
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"general_email", fields(generalUser), "general@example.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &GeneralUser{
				UserUUID: tt.fields.UserUUID,
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
				ClubUUID: tt.fields.ClubUUID,
			}
			got := u.GetEmail()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetEmail() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestGeneralUser_GetName(t *testing.T) {
	type fields struct {
		UserUUID string
		Email    string
		Name     string
		ClubUUID sql.NullString
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"general_name", fields(generalUser), "general example"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &GeneralUser{
				UserUUID: tt.fields.UserUUID,
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
				ClubUUID: tt.fields.ClubUUID,
			}
			got := u.GetName()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetName() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestGeneralUser_GetRole(t *testing.T) {
	type fields struct {
		UserUUID string
		Email    string
		Name     string
		ClubUUID sql.NullString
	}
	tests := []struct {
		name   string
		fields fields
		want   consts.UserType
	}{
		{"general_role", fields(generalUser), "general"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &GeneralUser{
				UserUUID: tt.fields.UserUUID,
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
				ClubUUID: tt.fields.ClubUUID,
			}
			got := u.GetRole()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetRole() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestGeneralUser_GetUserID(t *testing.T) {
	type fields struct {
		UserUUID string
		Email    string
		Name     string
		ClubUUID sql.NullString
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"general_uuid", fields(generalUser), consts.DummyUUID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &GeneralUser{
				UserUUID: tt.fields.UserUUID,
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
				ClubUUID: tt.fields.ClubUUID,
			}
			got := u.GetUserID()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetUserID() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestGeneralUser_ToUserResponse(t *testing.T) {
	type fields struct {
		UserUUID string
		Email    string
		Name     string
		ClubUUID sql.NullString
	}
	tests := []struct {
		name   string
		fields fields
		want   *UserResponse
	}{
		{"general_response", fields(generalUser), &UserResponse{consts.DummyUUID, "general@example.com", "general example", "general"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &GeneralUser{
				UserUUID: tt.fields.UserUUID,
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
				ClubUUID: tt.fields.ClubUUID,
			}
			got := u.ToUserResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToUserResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
