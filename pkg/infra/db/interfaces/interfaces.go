package interfaces

import (
	"gorm.io/gorm"
)

type (
	// Connection is an interface for database connection operations.
	Connection interface {
		// GetRawDBConnection returns an instance of gorm.DB which represents a raw database connection.
		GetRawDBConnection() *gorm.DB
	}
)
