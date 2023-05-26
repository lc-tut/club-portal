package admins

import (
	"github.com/lc-tut/club-portal/repos/clubs"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestNewAdminRepository(t *testing.T) {
	type args struct {
		logger   *zap.Logger
		db       *gorm.DB
		clubRepo clubs.IClubRepository
	}
	tests := []struct {
		name string
		args args
		want *AdminRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdminRepository(tt.args.logger, tt.args.db, tt.args.clubRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdminRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
