package repository

import (
	"database/sql"
	"log"

	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (r *Repository) ExtractAmount(name string) (float32, error) {
	row := r.db.QueryRow(ExtractNameAmountQuery, name)

	var value entity.Value

	if err := row.Scan(&value.Amount); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		} else {
			log.Printf("Error al escanear", err)
			return 0, err
		}

	}
	return value.Amount, nil
}
