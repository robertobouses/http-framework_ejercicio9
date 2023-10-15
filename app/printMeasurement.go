package app

import (
	"log"

	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (s *Service) PrintMeasurement() ([]entity.Measurement, error) {
	measurement, err := s.repo.ExtractMeasurement()
	if err != nil {
		log.Printf("Error al extraer mediciones", err)
		return []entity.Measurement{}, err
	}
	return measurement, nil

}
