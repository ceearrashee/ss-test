package store

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	_tableNameCache sync.Map
)

func getObjectSchema[T any](db *gorm.DB, t T) (*schema.Schema, error) {
	sch, err := schema.Parse(t, &_tableNameCache, db.NamingStrategy)
	if err != nil {
		return nil, fmt.Errorf("error getting object schema: %w", err)
	}

	return sch, nil
}

func getFieldDetails[T any](db *gorm.DB, t T, columnName string) (*schema.Field, error) {
	sch, err := getObjectSchema(db, t)
	if err != nil {
		return nil, fmt.Errorf("error getting field details for column: %s, error: %w", columnName, err)
	}

	field := sch.LookUpField(columnName)

	return field, nil
}

// GetObjectDBFieldName function attempts to fetch a DB field name for an object based on the specified column name.
// Returns the field name if successful, and an error if not.
func GetObjectDBFieldName[T any](db *gorm.DB, t T, columnName string) (string, error) {
	field, err := getFieldDetails(db, t, columnName)
	if err != nil {
		return "", fmt.Errorf("error getting DB field name for column: %s, error: %w", columnName, err)
	}

	return field.DBName, nil
}

// GetObjectFieldName function attempts to fetch an object field name based on the specified column name.
// Returns the field name if successful, and an error if not.
func GetObjectFieldName[T any](db *gorm.DB, t T, columnName string) (string, error) {
	field, err := getFieldDetails(db, t, columnName)
	if err != nil {
		return "", fmt.Errorf("error getting object field name for column: %s, error: %w", columnName, err)
	}

	return field.Name, nil
}

// GetTableName attempts to fetch the table name for an object.
// Returns the table name if successful, and an error if not.
func GetTableName[T any](db *gorm.DB, t T) (string, error) {
	modelSchema, err := schema.Parse(t, &_tableNameCache, db.NamingStrategy)
	if err != nil {
		return "", fmt.Errorf("error parsing for table name, error: %w", err)
	}

	return modelSchema.Table, nil
}

// GetDBObjectField function returns DB name for a field.
// It generates table name then converts the field into DB field name.
// Returns the DB name if successful, and an error if not.
func GetDBObjectField[T any](db *gorm.DB, field string) (string, error) {
	var dbModel T

	tableName, err := GetTableName(db, dbModel)
	if err != nil {
		return "", fmt.Errorf("error getting table name for field: %s, error: %w", field, err)
	}

	fieldName, err := GetObjectDBFieldName(db, dbModel, field)
	if err != nil {
		return "", fmt.Errorf("error getting DB object field for field: %s, error: %w", field, err)
	}

	return db.NamingStrategy.ColumnName(tableName, fieldName), nil
}
