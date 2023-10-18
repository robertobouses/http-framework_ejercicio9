package app

import (
	"github.com/robertobouses/http-framework_ejercicio9/entity"
	"github.com/schollz/fuzzy"
)

func (s *Service) SearchNameValue(nameValueRequest entity.NameValueRequest) ([]string, error) {
	values, err := s.repo.ExtractValues()
	if err != nil {
		return nil, err
	}

	threshold := 0.6
	var results []string
	for _, value := range values {
		matches := fuzzy.Find(nameValueRequest.Name, []string{value.Name})
		if len(matches) > 0 && matches[0].Score >= threshold {
			results = append(results, value.Name)
		}
	}

	return results, nil
}
