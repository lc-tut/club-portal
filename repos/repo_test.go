package repos

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/lc-tut/club-portal/repos/admins"
	"github.com/lc-tut/club-portal/repos/clubs"
	"github.com/lc-tut/club-portal/repos/users"
	"github.com/lc-tut/club-portal/testutil"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var logger = testutil.NewTestZapLogger()
var db, _, _ = testutil.NewUnitTestMockDB()
var clubRepo = clubs.NewClubRepository(logger, db)
var userRepo = users.NewUserRepository(logger, db)
var adminRepo = admins.NewAdminRepository(logger, db, clubRepo)

func TestNewRepository(t *testing.T) {
	type args struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *Repository
	}{
		{
			"new_repo", args{logger, db}, &Repository{
				clubRepo,
				userRepo,
				adminRepo,
				logger,
				db,
			},
		},
	}
	opts := []cmp.Option{
		cmpopts.IgnoreTypes(&zap.Logger{}, &gorm.DB{}),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRepository(tt.args.logger, tt.args.db)
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("NewRepository() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
