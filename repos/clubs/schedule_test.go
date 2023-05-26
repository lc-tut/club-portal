package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestClubRepository_CreateSchedule(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubUUID string
		args     []ClubScheduleArgs
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
			if err := r.CreateSchedule(tt.args.clubUUID, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateSchedule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_CreateScheduleWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx       *gorm.DB
		clubUUID string
		args     []ClubScheduleArgs
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
			if err := r.CreateScheduleWithTx(tt.args.tx, tt.args.clubUUID, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateScheduleWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_GetScheduleByID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		scheduleID uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clubs.ClubSchedule
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
			got, err := r.GetScheduleByID(tt.args.scheduleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetScheduleByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetScheduleByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetSchedulesByClubUUID(t *testing.T) {
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
		want    []clubs.ClubSchedule
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
			got, err := r.GetSchedulesByClubUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSchedulesByClubUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSchedulesByClubUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_UpdateSchedule(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubUUID string
		args     []ClubScheduleArgs
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
			if err := r.UpdateSchedule(tt.args.clubUUID, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSchedule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_UpdateScheduleWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx       *gorm.DB
		clubUUID string
		args     []ClubScheduleArgs
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
			if err := r.UpdateScheduleWithTx(tt.args.tx, tt.args.clubUUID, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdateScheduleWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
