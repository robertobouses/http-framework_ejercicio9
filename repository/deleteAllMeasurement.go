package repository

func (r *Repository) DeleteAllMeasurement() error {
	_, err := r.db.Exec(DeleteAllMeasurementQuery)
	return err
}
