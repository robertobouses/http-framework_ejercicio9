package repository

import (
	"database/sql"

	_ "embed"

	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

type REPOSITORY interface {
	InsertMeasurement(medicion entity.Measurement) error
	ExtractMeasurement() ([]entity.Measurement, error)
	DeleteMeasurement(id string) error
	ExtractMeasurementId(id string) (entity.Measurement, error)
	DeleteAllMeasurement() error
	DeleteEmptyMeasurement() error
	InsertValue(value entity.Value) error
	ExtractAmount(name string) (float32, error)
	ExtractValues() ([]entity.Value, error)
}

type Repository struct {
	db                    *sql.DB
	insertStmt            *sql.Stmt
	extractStmt           *sql.Stmt
	deleteStmt            *sql.Stmt
	extractIdStmt         *sql.Stmt
	deleteAllStmt         *sql.Stmt
	deleteEmptyStmt       *sql.Stmt
	insertValueStmt       *sql.Stmt
	extractNameAmountStmt *sql.Stmt
	extractValuesStmt     *sql.Stmt
}

//go:embed sql/insert_measurement.sql
var InsertMeasurementQuery string

//go:embed sql/extract_measurement.sql
var ExtractMeasurementQuery string

//go:embed sql/delete_measurement.sql
var DeleteMeasurementQuery string

//go:embed sql/extractid_measurement.sql
var ExtractIdMeasurementQuery string

//go:embed sql/deleteall_measurement.sql
var DeleteAllMeasurementQuery string

//go:embed sql/deleteempty_measurement.sql
var DeleteEmptyMeasurementQuery string

//go:embed sql/insert_value.sql
var InsertValueQuery string

//go:embed sql/extractname_amount.sql
var ExtractNameAmountQuery string

//go:embed sql/extract_values.sql
var ExtractValuesQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	insertStmt, err := db.Prepare(InsertMeasurementQuery)
	if err != nil {
		return nil, err
	}

	extractStmt, err := db.Prepare(ExtractMeasurementQuery)
	if err != nil {
		return nil, err
	}

	deleteStmt, err := db.Prepare(DeleteMeasurementQuery)
	if err != nil {
		return nil, err
	}

	extractIdStmt, err := db.Prepare(ExtractIdMeasurementQuery)
	if err != nil {
		return nil, err
	}

	deleteAllStmt, err := db.Prepare(DeleteAllMeasurementQuery)
	if err != nil {
		return nil, err
	}

	deleteEmptyStmt, err := db.Prepare(DeleteEmptyMeasurementQuery)
	if err != nil {
		return nil, err
	}

	insertValueStmt, err := db.Prepare(InsertValueQuery)
	if err != nil {
		return nil, err
	}

	extractNameAmountStmt, err := db.Prepare(ExtractNameAmountQuery)
	if err != nil {
		return nil, err
	}
	extractValuesStmt, err := db.Prepare(ExtractValuesQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                    db,
		insertStmt:            insertStmt,
		extractStmt:           extractStmt,
		deleteStmt:            deleteStmt,
		extractIdStmt:         extractIdStmt,
		deleteAllStmt:         deleteAllStmt,
		deleteEmptyStmt:       deleteEmptyStmt,
		insertValueStmt:       insertValueStmt,
		extractNameAmountStmt: extractNameAmountStmt,
		extractValuesStmt:     extractValuesStmt,
	}, nil
}
