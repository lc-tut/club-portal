package admins

import (
	"github.com/lc-tut/club-portal/models/users"
	"github.com/lc-tut/club-portal/repos/clubs"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestAdminRepository_GetAllUser(t *testing.T) {
	type fields struct {
		logger          *zap.Logger
		db              *gorm.DB
		IClubRepository clubs.IClubRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.UserInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdminRepository{
				logger:          tt.fields.logger,
				db:              tt.fields.db,
				IClubRepository: tt.fields.IClubRepository,
			}
			got, err := r.GetAllUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminRepository_GetSpecifiedUser(t *testing.T) {
	type fields struct {
		logger          *zap.Logger
		db              *gorm.DB
		IClubRepository clubs.IClubRepository
	}
	type args struct {
		userUUID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    users.UserInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdminRepository{
				logger:          tt.fields.logger,
				db:              tt.fields.db,
				IClubRepository: tt.fields.IClubRepository,
			}
			got, err := r.GetSpecifiedUser(tt.args.userUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpecifiedUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSpecifiedUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminRepository_UpdateSpecifiedDomainUser(t *testing.T) {
	type fields struct {
		logger          *zap.Logger
		db              *gorm.DB
		IClubRepository clubs.IClubRepository
	}
	type args struct {
		userUUID string
		name     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdminRepository{
				logger:          tt.fields.logger,
				db:              tt.fields.db,
				IClubRepository: tt.fields.IClubRepository,
			}
			if err := r.UpdateSpecifiedDomainUser(tt.args.userUUID, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSpecifiedDomainUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdminRepository_UpdateSpecifiedGeneralUser(t *testing.T) {
	type fields struct {
		logger          *zap.Logger
		db              *gorm.DB
		IClubRepository clubs.IClubRepository
	}
	type args struct {
		userUUID string
		args     UserArgs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdminRepository{
				logger:          tt.fields.logger,
				db:              tt.fields.db,
				IClubRepository: tt.fields.IClubRepository,
			}
			if err := r.UpdateSpecifiedGeneralUser(tt.args.userUUID, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSpecifiedGeneralUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdminRepository_getAdminUser(t *testing.T) {
	type fields struct {
		logger          *zap.Logger
		db              *gorm.DB
		IClubRepository clubs.IClubRepository
	}
	type args struct {
		userUUID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.AdminUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdminRepository{
				logger:          tt.fields.logger,
				db:              tt.fields.db,
				IClubRepository: tt.fields.IClubRepository,
			}
			got, err := r.getAdminUser(tt.args.userUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAdminUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAdminUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminRepository_getDomainUser(t *testing.T) {
	type fields struct {
		logger          *zap.Logger
		db              *gorm.DB
		IClubRepository clubs.IClubRepository
	}
	type args struct {
		userUUID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.DomainUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdminRepository{
				logger:          tt.fields.logger,
				db:              tt.fields.db,
				IClubRepository: tt.fields.IClubRepository,
			}
			got, err := r.getDomainUser(tt.args.userUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDomainUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDomainUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminRepository_getGeneralUser(t *testing.T) {
	type fields struct {
		logger          *zap.Logger
		db              *gorm.DB
		IClubRepository clubs.IClubRepository
	}
	type args struct {
		userUUID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.GeneralUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AdminRepository{
				logger:          tt.fields.logger,
				db:              tt.fields.db,
				IClubRepository: tt.fields.IClubRepository,
			}
			got, err := r.getGeneralUser(tt.args.userUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getGeneralUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGeneralUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
