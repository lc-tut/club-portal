package repos

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.logger, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
