package testutil

import (
	"github.com/DATA-DOG/go-sqlmock"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewTestZapLogger() *zap.Logger {
	return zap.NewExample()
}

func NewUnitTestMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{Conn: mockDB}))
	if err != nil {
		return nil, nil, err
	}
	return db, mock, nil
}
