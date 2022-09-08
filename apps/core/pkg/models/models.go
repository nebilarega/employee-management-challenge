package models

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Model facilitate database interactions
type Model struct {
	models map[string]reflect.Value
	isOpen bool
	*gorm.DB
}

// NewModel returns a new Model without opening database connection
func NewModel() *Model {
	return &Model{
		models: make(map[string]reflect.Value),
	}
}

// IsOpen returns true if the Model has already established connection
// to the database
func (m *Model) IsOpen() bool {
	return m.isOpen
}

// OpenPostgres ...
func (m *Model) OpenPostgres() error {
	dbHost := os.Getenv("DB_HOST")
	// dbReplicationHost := os.Getenv("DB_REPLICATION_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable TimeZone=UTC password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		return err
	}

	m.DB = db
	m.isOpen = true

	return nil
}

// Register adds the values to the models registry
func (m *Model) Register(values ...interface{}) error {

	// do not work on them.models first, this is like an insurance policy
	// whenever we encounter any error in the values nothing goes into the registry
	models := make(map[string]reflect.Value)
	if len(values) > 0 {
		for _, val := range values {
			rVal := reflect.ValueOf(val)
			if rVal.Kind() == reflect.Ptr {
				rVal = rVal.Elem()
			}
			switch rVal.Kind() {
			case reflect.Struct:
				models[getTypeName(rVal.Type())] = reflect.New(rVal.Type())
			default:
				return errors.New("models must be structs")
			}
		}
	}
	for k, v := range models {
		m.models[k] = v
	}
	return nil
}

// AutoMigrateAll runs migrations for all the registered models
func (m *Model) AutoMigrateAll() {
	for _, v := range m.models {
		m.AutoMigrate(v.Interface())
	}
}

// DropAll drops all tables
func (m *Model) DropAll() {
	for _, v := range m.models {
		m.Migrator().DropTable(v.Interface())
	}
}

// RegisterAllModels ...
func (m *Model) RegisterAllModels() {
	m.Register(Department{})
	m.Register(Employee{})
}

func getTypeName(typ reflect.Type) string {
	if typ.Name() != "" {
		return typ.Name()
	}
	split := strings.Split(typ.String(), ".")
	return split[len(split)-1]
}
