package admins

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/lc-tut/club-portal/repos/clubs"
	"github.com/lc-tut/club-portal/testutil"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var logger = testutil.NewTestZapLogger()
var db, _, _ = testutil.NewUnitTestMockDB()
var clubRepo = clubs.NewClubRepository(logger, db)

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
		{"new_repo_admin", args{
			logger:   logger,
			db:       db,
			clubRepo: clubRepo,
		},
			&AdminRepository{
				logger:          logger,
				db:              db,
				IClubRepository: clubRepo,
			},
		},
	}
	opts := []cmp.Option{
		cmpopts.IgnoreTypes(&zap.Logger{}, &gorm.DB{}),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAdminRepository(tt.args.logger, tt.args.db, tt.args.clubRepo)
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("NewAdminRepository() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
