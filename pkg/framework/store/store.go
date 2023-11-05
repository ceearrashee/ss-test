package store

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type (
	// Condition represents a single filter condition.
	Condition struct {
		Column   string
		Operator string
		Value    any
	}

	// Filter is a function that transforms a Schema to a Condition.
	Filter func(ctx context.Context, sch schema.Schema) Condition

	// Repository is an interface that defines a standard set of CRUD operations.
	Repository[TDBModel any] interface {
		Save(ctx context.Context, obj *TDBModel) error
		GetByID(ctx context.Context, id uint) (*TDBModel, error)
		RecordExistsByID(ctx context.Context, id uint) (bool, error)
		GetWithFilter(ctx context.Context, filters ...Filter) ([]TDBModel, error)
		DeleteByID(ctx context.Context, id uint) error
	}

	// FromEntityFN is a function that converts an Entity to a DBModel.
	FromEntityFN[TEntity, TDBModel any] func(context.Context, *TEntity) (*TDBModel, error)

	// ToEntityFN is a function that converts a DBModel to an Entity.
	ToEntityFN[TEntity, TDBModel any] func(context.Context, *TDBModel) (*TEntity, error)

	// BaseStore is a structure that provides a basic implementation of the Repository interface.
	BaseStore[TEntity, TDBModel any] struct {
		DB         *gorm.DB
		FromEntity FromEntityFN[TEntity, TDBModel]
		ToEntity   ToEntityFN[TEntity, TDBModel]
	}
)

// New creates a new Store with specified conversion functions and DB connection.
func New[TEntity, TDBModel any](
	db *gorm.DB,
	fromEntity FromEntityFN[TEntity, TDBModel],
	toEntity ToEntityFN[TEntity, TDBModel],
) *BaseStore[TEntity, TDBModel] {
	return &BaseStore[TEntity, TDBModel]{
		DB:         db,
		FromEntity: fromEntity,
		ToEntity:   toEntity,
	}
}

// Save stores the entity into the database.
func (s *BaseStore[TEntity, TDBModel]) Save(ctx context.Context, entity *TEntity) error {
	dbModel, err := s.FromEntity(ctx, entity)
	if err != nil {
		return fmt.Errorf("converting entity to DB model: %w", err)
	}

	err = s.DB.WithContext(ctx).Save(dbModel).Error
	if err != nil {
		return fmt.Errorf("saving DB model: %w", err)
	}

	newEntity, err := s.ToEntity(ctx, dbModel)
	if err != nil {
		return fmt.Errorf("converting DB model to entity: %w", err)
	}

	*entity = *newEntity

	return nil
}

// GetByID retrieves an entity given its ID.
func (s *BaseStore[TEntity, TDBModel]) GetByID(ctx context.Context, entityID uint) (*TEntity, error) {
	var dbModel TDBModel

	const idColumn = "ID"

	objectIDFieldName, err := GetDBObjectField[TDBModel](s.DB, idColumn)
	if err != nil {
		return nil, fmt.Errorf("retrieving object ID field name %q: %w", idColumn, err)
	}

	err = s.DB.WithContext(ctx).First(&dbModel, fmt.Sprintf("%s = ?", objectIDFieldName), entityID).Error
	if err != nil {
		return nil, fmt.Errorf("retrieving DB model by ID %d: %w", entityID, err)
	}

	entity, err := s.ToEntity(ctx, &dbModel)
	if err != nil {
		return nil, fmt.Errorf("converting DB model with ID %d to entity: %w", entityID, err)
	}

	return entity, nil
}

// RecordExistsByID checks whether a record exists with the given ID.
func (s *BaseStore[TEntity, TDBModel]) RecordExistsByID(ctx context.Context, entityID uint) (bool, error) {
	var dbModel TDBModel

	const idColumn = "ID"

	objectIDFieldName, err := GetDBObjectField[TDBModel](s.DB, idColumn)
	if err != nil {
		return false, fmt.Errorf("retrieving object ID field name %q: %w", idColumn, err)
	}

	err = s.DB.WithContext(ctx).Select("null").First(&dbModel, fmt.Sprintf("%s = ?", objectIDFieldName), entityID).Error

	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return false, nil
	default:
		return false, fmt.Errorf("checking if record exists by ID %d: %w", entityID, err)
	}
}

// GetWithFilter retrieves records from the database that meet the specified filter conditions.
func (s *BaseStore[TEntity, TDBModel]) GetWithFilter(
	ctx context.Context,
	filters ...Filter,
) ([]TEntity, error) {
	var (
		dbModels []TDBModel
		dbModel  TDBModel
	)

	objectSchema, err := getObjectSchema(s.DB, dbModel)
	if err != nil {
		return nil, fmt.Errorf("getting object schema: %w", err)
	}

	qrs := make([]any, 0, len(filters))

	for _, filter := range filters {
		condition := filter(ctx, *objectSchema)
		qrs = append(qrs, gorm.Expr(fmt.Sprintf("%s %s ?", condition.Column, condition.Operator), condition.Value))
	}

	err = s.DB.WithContext(ctx).Find(&dbModels, qrs...).Error
	if err != nil {
		return nil, fmt.Errorf("finding DB models with filter: %w", err)
	}

	entities := make([]TEntity, 0, len(dbModels))

	for i := range dbModels {
		model := dbModels[i]

		entity, err := s.ToEntity(ctx, &model)
		if err != nil {
			return nil, fmt.Errorf("converting DB model to entity: %w", err)
		}

		entities = append(entities, *entity)
	}

	return entities, nil
}

// DeleteByID deletes an entity represented by the given ID from the database.
func (s *BaseStore[TEntity, TDBModel]) DeleteByID(ctx context.Context, entityID uint) error {
	var dbModel TDBModel

	const idColumn = "ID"

	objectIDFieldName, err := GetDBObjectField[TDBModel](s.DB, idColumn)
	if err != nil {
		return fmt.Errorf("retrieving object ID field name %q: %w", idColumn, err)
	}

	err = s.DB.WithContext(ctx).Delete(&dbModel, fmt.Sprintf("%s = ?", objectIDFieldName), entityID).Error
	if err != nil {
		return fmt.Errorf("deleting DB model by ID: %w", err)
	}

	return nil
}
