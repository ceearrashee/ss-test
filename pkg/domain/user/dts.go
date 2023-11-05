package user

import (
	"context"
	"time"

	"gorm.io/gorm"

	"solid-software.test-task/pkg/infra/db/models"
)

type (
	// Entity represents a user domain model.
	// It is used to transfer data between the domain and the infrastructure layers.
	Entity struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Name      string    `json:"name,omitempty"`
		Surname   string    `json:"surname,omitempty"`
		Phone     string    `json:"phone,omitempty"`
		Address   string    `json:"address,omitempty"`
	}
)

func toEntity(_ context.Context, dbUser *models.User) (*Entity, error) {
	entityUser := Entity{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		Surname:   dbUser.Surname,
		Phone:     dbUser.Phone,
		Address:   dbUser.Address,
	}

	return &entityUser, nil
}

func toDBModel(_ context.Context, entity *Entity) (*models.User, error) {
	dbModel := models.User{
		Model: gorm.Model{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
		Name:    entity.Name,
		Surname: entity.Surname,
		Phone:   entity.Phone,
		Address: entity.Address,
	}

	return &dbModel, nil
}
