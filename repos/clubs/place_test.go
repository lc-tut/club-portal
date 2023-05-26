package clubs

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestClubRepository_CreatePlace(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		args []ClubPlaceArgs
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
			if err := r.CreatePlace(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreatePlace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_CreatePlaceWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx   *gorm.DB
		args []ClubPlaceArgs
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
			if err := r.CreatePlaceWithTx(tt.args.tx, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreatePlaceWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_GetPlaceByID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		placeID uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clubs.ClubPlace
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
			got, err := r.GetPlaceByID(tt.args.placeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlaceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlaceByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetPlacesByClubUUID(t *testing.T) {
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
		want    []clubs.ClubPlace
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
			got, err := r.GetPlacesByClubUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlacesByClubUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlacesByClubUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
