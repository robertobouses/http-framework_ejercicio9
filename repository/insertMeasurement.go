package repository

import "github.com/robertobouses/http-framework_ejercicio9/entity"

func (r *Repository) InsertMeasurement(medicion entity.Measurement) error {
	_, err := r.insertStmt.Exec(medicion.Id, medicion.ValorX, medicion.ValorY, medicion.ValorZ)
	return err
}
