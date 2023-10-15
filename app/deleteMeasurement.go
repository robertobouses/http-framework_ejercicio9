package app

import "log"

func (s *Service) DeleteMeasurement(id string) error {
	err := s.repo.DeleteMeasurement(id)
	if err != nil {
		log.Printf("Error", err)
		return err
	}
	return nil
}
