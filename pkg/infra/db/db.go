package db

import (
	"gorm.io/gorm"

	"solid-software.test-task/pkg/infra/db/di"
)

var (
	// ErrRecordNotFound is a wrapper for gorm.ErrRecordNotFound.
	//  It is used to avoid importing gorm in domain layer.
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

// GetRawDBConnection returns a raw DB connection.
func GetRawDBConnection() *gorm.DB {
	return di.InitializeNewDBConnection().GetRawDBConnection()
}
