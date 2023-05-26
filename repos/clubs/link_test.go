package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestClubRepository_CreateLink(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubUUID string
		args     []ClubLinkArgs
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
			if err := r.CreateLink(tt.args.clubUUID, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateLink() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_CreateLinkWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx       *gorm.DB
		clubUUID string
		args     []ClubLinkArgs
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
			if err := r.CreateLinkWithTx(tt.args.tx, tt.args.clubUUID, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateLinkWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_GetAllLinks(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []clubs.ClubLink
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
			got, err := r.GetAllLinks()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllLinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllLinks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetLinkByID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		linkID uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clubs.ClubLink
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
			got, err := r.GetLinkByID(tt.args.linkID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLinkByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLinkByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetLinksByClubUUID(t *testing.T) {
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
		want    []clubs.ClubLink
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
			got, err := r.GetLinksByClubUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLinksByClubUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLinksByClubUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_UpdateLink(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubUUID string
		args     []ClubLinkArgs
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
			if err := r.UpdateLink(tt.args.clubUUID, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLink() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_UpdateLinkWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx       *gorm.DB
		clubUUID string
		args     []ClubLinkArgs
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
			if err := r.UpdateLinkWithTx(tt.args.tx, tt.args.clubUUID, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLinkWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
