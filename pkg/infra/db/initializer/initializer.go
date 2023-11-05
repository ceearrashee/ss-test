package initializer

import (
	"fmt"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"solid-software.test-task/pkg/infra/db/interfaces"
	"solid-software.test-task/pkg/infra/db/models"
)

type (
	connectionImpl struct {
		*gorm.DB
	}
)

var (
	_dbInitOnce   sync.Once //nolint:gochecknoglobals
	_dbConnection *gorm.DB  //nolint:gochecknoglobals
)

// GetRawDBConnection returns a raw DB connection.
func (c connectionImpl) GetRawDBConnection() *gorm.DB {
	switch {
	case c.DB != nil:
		return c.DB
	case _dbConnection != nil:
		return _dbConnection
	default:
		panic("db connection is not initialized")
	}
}

// InitDBConnection initializes db connection.
func InitDBConnection() interfaces.Connection {
	_dbInitOnce.Do(
		func() {
			// TODO: implement different db types and add config for it.
			dbConnection, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
			if err != nil {
				panic(fmt.Errorf("open db connection: %w", err))
			}

			_dbConnection = dbConnection
			// TODO: for right db migration must be used github.com/pressly/goose or something like this.
			//  but for this test task it's not necessary
			err = migrateDBModels(_dbConnection, &models.User{})
			if err != nil {
				panic(err)
			}
		},
	)

	return &connectionImpl{_dbConnection}
}

func migrateDBModels(db *gorm.DB, dbModels ...any) error {
	if err := db.Migrator().AutoMigrate(dbModels...); err != nil {
		return fmt.Errorf("migrate db models: %w", err)
	}

	return nil
}
