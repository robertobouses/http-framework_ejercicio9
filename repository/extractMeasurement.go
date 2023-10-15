package repository

import (
	"log"

	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (r *Repository) ExtractMeasurement() ([]entity.Measurement, error) {
	log.Println("SQL Query", ExtractMeasurementQuery)
	rows, err := r.db.Query(ExtractMeasurementQuery)

	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL", err)
		return nil, err
	}
	defer rows.Close()
	var measurements []entity.Measurement

	for rows.Next() {
		var measurement entity.Measurement
		if err := rows.Scan(&measurement.Id, &measurement.ValorX, &measurement.ValorY, &measurement.ValorZ); err != nil {
			log.Printf("Error al escanear filas", err)
			return nil, err
		}
		measurements = append(measurements, measurement)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al recorrer filas", err)
		return nil, err
	}

	return measurements, nil
}
