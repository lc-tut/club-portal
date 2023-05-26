package clubs

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestClubRepository_GetClubThumbnailByID(t *testing.T) {
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
		want    *clubs.ClubThumbnail
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ClubRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetClubThumbnailByID(tt.args.thumbnailID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClubThumbnailByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClubThumbnailByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetClubThumbnailByUUID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubUUID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clubs.ClubThumbnail
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ClubRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := r.GetClubThumbnailByUUID(tt.args.clubUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClubThumbnailByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClubThumbnailByUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_UpdateClubThumbnail(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubUUID    string
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
			r := &ClubRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			if err := r.UpdateClubThumbnail(tt.args.clubUUID, tt.args.thumbnailID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateClubThumbnail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_UpdateClubThumbnailWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx          *gorm.DB
		clubUUID    string
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
			r := &ClubRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			if err := r.UpdateClubThumbnailWithTx(tt.args.tx, tt.args.clubUUID, tt.args.thumbnailID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateClubThumbnailWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
