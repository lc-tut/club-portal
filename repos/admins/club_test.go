package admins

import (
	"github.com/lc-tut/club-portal/repos/clubs"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"testing"
)

func TestAdminRepository_UpdatePageByClubUUIDWithAdmin(t *testing.T) {
	type fields struct {
		logger          *zap.Logger
		db              *gorm.DB
		IClubRepository clubs.IClubRepository
	}
	type args struct {
		uuid string
		args ClubPageUpdateArgsWithAdmin
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
			r := &AdminRepository{
				logger:          tt.fields.logger,
				db:              tt.fields.db,
				IClubRepository: tt.fields.IClubRepository,
			}
			if err := r.UpdatePageByClubUUIDWithAdmin(tt.args.uuid, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePageByClubUUIDWithAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
