package users

import (
	"github.com/lc-tut/club-portal/models/users"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestUserRepository_CreateAdminUser(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid  string
		email string
		name  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.AdminUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.CreateAdminUser(tt.args.uuid, tt.args.email, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAdminUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAdminUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_CreateDomainUser(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid  string
		email string
		name  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.DomainUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.CreateDomainUser(tt.args.uuid, tt.args.email, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDomainUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDomainUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_CreateGeneralUser(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid  string
		email string
		name  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.GeneralUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.CreateGeneralUser(tt.args.uuid, tt.args.email, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGeneralUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGeneralUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetAdminUserByEmail(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.AdminUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetAdminUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdminUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdminUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetAdminUserByUUID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.AdminUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetAdminUserByUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdminUserByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdminUserByUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetAllGeneralUser(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []users.GeneralUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetAllGeneralUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllGeneralUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllGeneralUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetDomainUserByEmail(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.DomainUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetDomainUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDomainUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDomainUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetDomainUserByUUID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.DomainUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetDomainUserByUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDomainUserByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDomainUserByUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetGeneralUserByEmail(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.GeneralUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetGeneralUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGeneralUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGeneralUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetGeneralUserByUUID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.GeneralUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetGeneralUserByUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGeneralUserByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGeneralUserByUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetUserByEmailFromRole(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		email string
		role  string
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
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetUserByEmailFromRole(tt.args.email, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmailFromRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmailFromRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetUserByUUIDFromRole(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
		role string
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
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetUserByUUIDFromRole(tt.args.uuid, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUUIDFromRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByUUIDFromRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_UpdateAdminUser(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
		name string
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
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			if err := r.UpdateAdminUser(tt.args.uuid, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdminUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_UpdateDomainUser(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
		name string
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
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			if err := r.UpdateDomainUser(tt.args.uuid, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDomainUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_UpdateGeneralUser(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid     string
		name     string
		clubUUID string
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
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			if err := r.UpdateGeneralUser(tt.args.uuid, tt.args.name, tt.args.clubUUID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateGeneralUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_UpdateUserFromRole(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
		role string
		args UpdateUserArgs
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
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			if err := r.UpdateUserFromRole(tt.args.uuid, tt.args.role, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserFromRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_createUser(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx   *gorm.DB
		uuid string
		role string
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
			r := &UserRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			if err := r.createUser(tt.args.tx, tt.args.uuid, tt.args.role); (err != nil) != tt.wantErr {
				t.Errorf("createUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
