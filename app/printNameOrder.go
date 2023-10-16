package app

import (
	"log"
	"sort"
)

func (s *Service) PrintNameOrder() ([]string, error) {
	values, err := s.repo.ExtractValues()
	if err != nil {
		log.Printf("Error al extraer valores", err)
		return nil, err
	}

	type ValueByAmount struct {
		Name   string
		Amount float32
	}

	var valuesByAmount []ValueByAmount

	for _, v := range values {
		valuesByAmount = append(valuesByAmount, ValueByAmount{
			Name:   v.Name,
			Amount: v.Amount,
		})
	}

	sort.Slice(valuesByAmount, func(i, j int) bool {
		return valuesByAmount[i].Amount < valuesByAmount[j].Amount
	})

	var orderedNames []string
	for _, v := range valuesByAmount {
		orderedNames = append(orderedNames, v.Name)
	}

	return orderedNames, nil
}
