package repository

import (
	"log"

	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (r *Repository) ExtractValues() ([]entity.Value, error) {
	rows, err := r.db.Query(ExtractValuesQuery)

	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL", err)
		return nil, err
	}
	defer rows.Close()

	var values []entity.Value

	for rows.Next() {
		var value entity.Value
		if err := rows.Scan(&value.Name, &value.Amount, &value.Invoiced); err != nil {
			log.Printf("Error al escanear filas", err)
			return nil, err
		}
		values = append(values, value)

	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al recorrer filas")
		return nil, err
	}
	return values, nil
}
