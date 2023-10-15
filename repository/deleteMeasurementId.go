package repository

func (r *Repository) DeleteMeasurement(id string) error {
	_, err := r.db.Exec(DeleteMeasurementQuery, id)
	return err
}
