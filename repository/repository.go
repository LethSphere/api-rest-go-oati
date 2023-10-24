package repository

import (
	"context"

	"vahar.com/go/rest-oati/models"
)

type Repository interface {
	InsertTutorial(cxt context.Context, tutorial *models.Tutorial) error
	GetTutorialById(cxt context.Context, id string) (*models.Tutorial, error)
	GetTutorialByTitle(cxt context.Context, titulo string) (*models.Tutorial, error)
	InsertDetail(cxt context.Context, detalle *models.Detalle) error
	Close() error
}

var implementations Repository

func SetRepository(repository Repository) {
	implementations = repository
}

// Agregar tutoriales
func InsertTutorial(cxt context.Context, tutorial *models.Tutorial) error {
	return implementations.InsertTutorial(cxt, tutorial)
}

// Buscar por Id
func GetTutorialById(cxt context.Context, id string) (*models.Tutorial, error) {
	return implementations.GetTutorialById(cxt, id)
}

// Buscar por titulo
func GetTutorialByTitle(cxt context.Context, titulo string) (*models.Tutorial, error) {
	return implementations.GetTutorialByTitle(cxt, titulo)
}

// Agregar detalle
func InsertDetail(cxt context.Context, detalle *models.Detalle) error {
	return implementations.InsertDetail(cxt, detalle)
}

func Close() error {
	return implementations.Close()
}
