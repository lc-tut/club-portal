package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestClubRepository_CreatePage(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
		args ClubPageCreateArgs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clubs.ClubPage
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
			got, err := r.CreatePage(tt.args.uuid, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_DeletePageByClubSlug(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		slug string
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
			if err := r.DeletePageByClubSlug(tt.args.slug); (err != nil) != tt.wantErr {
				t.Errorf("DeletePageByClubSlug() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_DeletePageByClubUUID(t *testing.T) {
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
			if err := r.DeletePageByClubUUID(tt.args.uuid); (err != nil) != tt.wantErr {
				t.Errorf("DeletePageByClubUUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_GetAllPages(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []clubs.ClubPageExternalInfo
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
			got, err := r.GetAllPages()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllPages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetPageByClubSlug(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubSlug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clubs.ClubPageInternalInfo
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
			got, err := r.GetPageByClubSlug(tt.args.clubSlug)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPageByClubSlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPageByClubSlug() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetPageByClubUUID(t *testing.T) {
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
		want    *clubs.ClubPageInternalInfo
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
			got, err := r.GetPageByClubUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPageByClubUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPageByClubUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetRestrictedPageByClubSlug(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubSlug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clubs.ClubPageRestrictedInfo
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
			got, err := r.GetRestrictedPageByClubSlug(tt.args.clubSlug)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRestrictedPageByClubSlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRestrictedPageByClubSlug() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_GetRestrictedPageByClubUUID(t *testing.T) {
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
		want    *clubs.ClubPageRestrictedInfo
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
			got, err := r.GetRestrictedPageByClubUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRestrictedPageByClubUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRestrictedPageByClubUUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_UpdatePageByClubSlug(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		clubSlug string
		args     ClubPageUpdateArgs
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
			if err := r.UpdatePageByClubSlug(tt.args.clubSlug, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePageByClubSlug() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_UpdatePageByClubUUID(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		uuid string
		args ClubPageUpdateArgs
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
			if err := r.UpdatePageByClubUUID(tt.args.uuid, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePageByClubUUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClubRepository_getPageInternal(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		page *clubs.ClubPage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clubs.ClubPageInternalInfo
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
			got, err := r.getPageInternal(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPageInternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPageInternal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClubRepository_getPageRestricted(t *testing.T) {
	type fields struct {
		logger *zap.Logger
		db     *gorm.DB
	}
	type args struct {
		page *clubs.ClubPage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clubs.ClubPageRestrictedInfo
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
			got, err := r.getPageRestricted(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPageRestricted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPageRestricted() got = %v, want %v", got, tt.want)
			}
		})
	}
}
