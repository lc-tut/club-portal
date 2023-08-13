package users

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestUserRepository_CreateFavorite(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		userUUID string
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
			if err := r.CreateFavorite(tt.args.userUUID, tt.args.clubUUID); (err != nil) != tt.wantErr {
				t.Errorf("CreateFavorite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_DeleteFavorite(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		userUUID string
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
			if err := r.DeleteFavorite(tt.args.userUUID, tt.args.clubUUID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFavorite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_GetFavorites(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		userUUID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []clubs.ClubPageExternalInfo
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
			got, err := r.GetFavorites(tt.args.userUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFavorites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFavorites() got = %v, want %v", got, tt.want)
			}
		})
	}
}
