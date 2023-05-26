package users

import (
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestUserRepository_CreateUploadedImage(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		userUUID string
		path     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.UploadedImage
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
			got, err := r.CreateUploadedImage(tt.args.userUUID, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUploadedImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUploadedImage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_DeleteImageByID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		imageID uint32
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
			if err := r.DeleteImageByID(tt.args.imageID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteImageByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_GetImagesByUserUUID(t *testing.T) {
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
		want    []users.UploadedImage
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
			got, err := r.GetImagesByUserUUID(tt.args.userUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImagesByUserUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImagesByUserUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetUploadedImageByID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		imageID uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *users.UploadedImage
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
			got, err := r.GetUploadedImageByID(tt.args.imageID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUploadedImageByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUploadedImageByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
