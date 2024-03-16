package users

import (
	"github.com/lc-tut/club-portal/models/users"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestUserRepository_CreateThumbnail(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.UploadedThumbnail
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
			got, err := r.CreateThumbnail(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateThumbnail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateThumbnail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_DeleteThumbnail(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		thumbnailID uint32
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
			if err := r.DeleteThumbnail(tt.args.thumbnailID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteThumbnail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_GetThumbnail(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		thumbnailID uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.UploadedThumbnail
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
			got, err := r.GetThumbnail(tt.args.thumbnailID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetThumbnail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetThumbnail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
