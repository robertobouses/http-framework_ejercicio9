package app

import (
	"fmt"

	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (s *Service) CreateMeasurement(newMeasurement entity.Measurement) error {

	existingMeasurement, err := s.repo.ExtractMeasurementId(newMeasurement.Id)
	if err != nil {

		return err
	}

	if existingMeasurement.Id != "" {

		return fmt.Errorf("El registro con ID %s ya existe", newMeasurement.Id)
	}

	return s.repo.InsertMeasurement(newMeasurement)
}
