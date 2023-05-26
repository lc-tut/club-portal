package clubs

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"testing"
)

func TestClubRepository_GetClubDescription(t *testing.T) {
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
		want    string
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
			got, err := r.GetClubDescription(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClubDescription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetClubDescription() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_UpdateClubDescription(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
		desc string
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
			if err := r.UpdateClubDescription(tt.args.uuid, tt.args.desc); (err != nil) != tt.wantErr {
				t.Errorf("UpdateClubDescription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
