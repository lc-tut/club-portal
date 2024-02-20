package users

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/lc-tut/club-portal/testutil"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var logger = testutil.NewTestZapLogger()
var db, _, _ = testutil.NewUnitTestMockDB()

func TestNewUserRepository(t *testing.T) {
	type args struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *UserRepository
	}{
		{"new_repo_user", args{
			logger: logger,
			db:     db,
		},
			&UserRepository{
				logger: logger,
				db:     db,
			},
		},
	}
	opts := []cmp.Option{
		cmpopts.IgnoreTypes(&zap.Logger{}, &gorm.DB{}),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserRepository(tt.args.logger, tt.args.db)
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("NewUserRepository() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
