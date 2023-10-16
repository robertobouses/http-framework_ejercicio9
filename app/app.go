package app

import (
	"github.com/robertobouses/http-framework_ejercicio9/entity"
	"github.com/robertobouses/http-framework_ejercicio9/repository"
)

type APP interface {
	CreateMeasurement(newMedicion entity.Measurement) error
	PrintMeasurement() ([]entity.Measurement, error)
	DeleteMeasurement(id string) error
	CalcCubic(id string) (int, error)
	FindMinAndMaxCubic(cubics []int) (int, int)
	DeleteAllMeasurement() error
	DeleteEmptyMeasurement() error
	CreateValue(entity.Value) error
	PrintAmount(name string) (float32, error)
}

type Service struct {
	repo repository.REPOSITORY
}

func NewAPP(repo repository.REPOSITORY) APP {
	return &Service{
		repo: repo,
	}
}
