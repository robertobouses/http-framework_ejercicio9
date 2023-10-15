package app

import "log"

func (s *Service) DeleteEmptyMeasurement() error {
	err := s.repo.DeleteEmptyMeasurement()
	if err != nil {
		log.Print("Error", err)
		return err
	}
	return nil
}
