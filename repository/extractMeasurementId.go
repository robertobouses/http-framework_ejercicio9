package repository

import (
	"database/sql"
	"log"

	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (r *Repository) ExtractMeasurementId(id string) (entity.Measurement, error) {
	//query := "SELECT id, valorx, valory, valorz FROM arithmetic.measurements WHERE id = $1"

	row := r.db.QueryRow(ExtractIdMeasurementQuery, id)

	var measurement entity.Measurement

	if err := row.Scan(&measurement.Id, &measurement.ValorX, &measurement.ValorY, &measurement.ValorZ); err != nil {
		if err == sql.ErrNoRows {
			return entity.Measurement{}, nil
		} else {
			log.Printf("Error al escanear fila", err)
			return entity.Measurement{}, err
		}
	}

	return measurement, nil
}
