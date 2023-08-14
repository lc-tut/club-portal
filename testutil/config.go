package testutil

import (
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewTestZapLogger() *zap.Logger {
	return zap.NewExample()
}

func NewUnitTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
		DryRun:                 true,
		DisableAutomaticPing:   true,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewIntegrationTestDB() (*gorm.DB, error) {
	// ? should use mysql driver?
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
		DisableAutomaticPing:   true,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
