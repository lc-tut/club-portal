package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestClubRepository_CreateTPRemark(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
		args []ClubTPRemarkArgs
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
			if err := r.CreateTPRemark(tt.args.uuid, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateTPRemark() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_CreateTPRemarkWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx   *gorm.DB
		uuid string
		args []ClubTPRemarkArgs
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
			if err := r.CreateTPRemarkWithTx(tt.args.tx, tt.args.uuid, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateTPRemarkWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_GetTPRemarksByClubUUID(t *testing.T) {
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
		want    []clubs.ClubRemark
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
			got, err := r.GetTPRemarksByClubUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTPRemarksByClubUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTPRemarksByClubUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_UpdateTPRemarkWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx   *gorm.DB
		uuid string
		args []ClubTPRemarkArgs
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
			if err := r.UpdateTPRemarkWithTx(tt.args.tx, tt.args.uuid, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTPRemarkWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
