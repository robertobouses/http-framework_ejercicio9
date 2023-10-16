package app

import "log"

func (s *Service) PrintValue(name string) (float32, error) {
	amount, err := s.repo.ExtractAmount(name)
	if err != nil {
		log.Printf("Error al obtener el VALUE por NAME", err)
		return 0, err
	}
	return amount, nil
}
