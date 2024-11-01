package clubs

import (
	"testing"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TestClubRepository_UpdateClubName(t *testing.T){
	type fieids struct{
		logger *zap.Logger
		db	   *gorm.DB
	}
	type args struct{
		uuid string
		name string
	}
	tests := []struct{
		name    string
		fieids  fieids
		args    args
		wantErr bool
	}{
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ClubRepository{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			if err := r.UpdateClubName(tt.args.uuid, tt.args.name); (err != nll) != tt.wantErr {
				t.Errorf("UpdateClubName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}