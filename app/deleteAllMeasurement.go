package app

import "log"

func (s *Service) DeleteAllMeasurement() error {
	err := s.repo.DeleteAllMeasurement()
	if err != nil {
		log.Print("Error", err)
		return err
	}
	return nil
}
