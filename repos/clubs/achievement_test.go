package clubs

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestClubRepository_CreateAchievement(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubUUID     string
		achievements []string
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
			if err := r.CreateAchievement(tt.args.clubUUID, tt.args.achievements); (err != nil) != tt.wantErr {
				t.Errorf("CreateAchievement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_CreateAchievementWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx           *gorm.DB
		clubUUID     string
		achievements []string
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
			if err := r.CreateAchievementWithTx(tt.args.tx, tt.args.clubUUID, tt.args.achievements); (err != nil) != tt.wantErr {
				t.Errorf("CreateAchievementWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_GetAchievementByID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		achievementID uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clubs.ClubAchievement
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
			got, err := r.GetAchievementByID(tt.args.achievementID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAchievementByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAchievementByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetAchievementsByClubUUID(t *testing.T) {
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
		want    []clubs.ClubAchievement
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
			got, err := r.GetAchievementsByClubUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAchievementsByClubUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAchievementsByClubUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_UpdateAchievement(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubUUID     string
		achievements []string
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
			if err := r.UpdateAchievement(tt.args.clubUUID, tt.args.achievements); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAchievement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_UpdateAchievementWithTx(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		tx           *gorm.DB
		clubUUID     string
		achievements []string
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
			if err := r.UpdateAchievementWithTx(tt.args.tx, tt.args.clubUUID, tt.args.achievements); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAchievementWithTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
