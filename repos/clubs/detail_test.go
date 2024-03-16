package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestClubRepository_CreateActivityDetail(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
		args []ActivityDetailArgs
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
			if err := r.CreateActivityDetail(tt.args.uuid, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateActivityDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_CreateActivityDetailWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx   *gorm.DB
		uuid string
		args []ActivityDetailArgs
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
			if err := r.CreateActivityDetailWithTx(tt.args.tx, tt.args.uuid, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateActivityDetailWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_GetActivityDetail(t *testing.T) {
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
		want    []clubs.ActivityDetail
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
			got, err := r.GetActivityDetail(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActivityDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActivityDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetAllRelations(t *testing.T) {
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
		want    []clubs.DetailRelations
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
			got, err := r.GetAllRelations(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllRelations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllRelations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_UpdateActivityDetailWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx   *gorm.DB
		uuid string
		args []ActivityDetailArgs
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
			if err := r.UpdateActivityDetailWithTx(tt.args.tx, tt.args.uuid, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdateActivityDetailWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_UpdateAllRelations(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid         string
		timeArgs     []ClubTimeArgs
		placeArgs    []ClubPlaceArgs
		detailArgs   []ActivityDetailArgs
		tpremarkArgs []ClubTPRemarkArgs
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
			if err := r.UpdateAllRelations(tt.args.uuid, tt.args.timeArgs, tt.args.placeArgs, tt.args.detailArgs, tt.args.tpremarkArgs); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAllRelations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
