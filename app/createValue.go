package app

import "github.com/robertobouses/http-framework_ejercicio9/entity"

func (s *Service) CreateValue(newValue entity.Value) error {
	return s.repo.InsertValue(newValue)
}
