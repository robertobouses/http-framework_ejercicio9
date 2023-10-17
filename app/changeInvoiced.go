package app

import "github.com/robertobouses/http-framework_ejercicio9/entity"

func (s *Service) ChangeInvoiced(amount float32) (entity.Value, error) {

	values, err := s.repo.ExtractValues()
	if err != nil {

		return entity.Value{}, err
	}

	for _, value := range values {
		if value.Amount == amount {
			value.Invoiced = !value.Invoiced
			valueChanged, err := s.repo.UpdateValue(value)
			if err != nil {
				return entity.Value{}, err

			}
			return valueChanged, nil
		}

	}
	return entity.Value{}, nil
}
