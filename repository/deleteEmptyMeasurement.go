package repository

func (r *Repository) DeleteEmptyMeasurement() error {
	_, err := r.db.Exec(DeleteEmptyMeasurementQuery)
	return err
}
