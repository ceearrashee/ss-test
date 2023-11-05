package user

import (
	"gorm.io/gorm"

	"solid-software.test-task/pkg/framework/store"
	"solid-software.test-task/pkg/infra/db/models"
)

type (
	// Service represents a user service with basic CRUD operations.
	Service interface {
		store.Repository[Entity]
	}

	service struct {
		*store.BaseStore[Entity, models.User]
	}
)

// NewUserService creates a new user service.
// It takes a database connection interface *gorm.DB as a parameter
// and returns an instance of UserEntity Service.
func NewUserService(db *gorm.DB) Service {
	s := new(service)
	s.BaseStore = store.New[Entity, models.User](db, toDBModel, toEntity)
	s.DB = db

	return s
}
