package repository

import "github.com/robertobouses/http-framework_ejercicio9/entity"

func (repo *Repository) UpdateValue(value entity.Value) (entity.Value, error) {

	updateQuery := "UPDATE arithmetic.values SET invoiced = $1 WHERE name = $2"

	_, err := repo.db.Exec(updateQuery, value.Invoiced, value.Name)

	if err != nil {
		return entity.Value{}, err
	}

	return value, nil
}
