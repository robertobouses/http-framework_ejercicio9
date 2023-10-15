package app

import "log"

func (s *Service) CalcCubic(id string) (int, error) {
	measurement, err := s.repo.ExtractMeasurementId(id)
	if err != nil {
		log.Printf("Error al obtener la medición por ID", err)
		return 0, err
	}
	cubic := measurement.ValorX * measurement.ValorY * measurement.ValorZ
	return cubic, nil
}
