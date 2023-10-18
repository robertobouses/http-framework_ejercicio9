package app

import (
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/robertobouses/http-framework_ejercicio9/entity"
)

func (s *Service) SearchNameValue(nameValueRequest entity.NameValueRequest) ([]string, error) {
	values, err := s.repo.ExtractValues()
	if err != nil {
		return nil, err
	}

	threshold := 0.6
	var results []string

	for _, value := range values {
		distance := fuzzy.Ratio(nameValueRequest.Name, value.Name)
		if distance >= threshold {
			results = append(results, value.Name)
		}
	}

	return results, nil
}
