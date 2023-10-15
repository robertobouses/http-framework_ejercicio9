package repository

import "github.com/robertobouses/http-framework_ejercicio9/entity"

func (r *Repository) InsertValue(value entity.Value) error {
	_, err := r.insertValueStmt.Exec(value.Name, value.Amount, value.Invoiced)
	return err
}
