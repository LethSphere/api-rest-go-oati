package repository

import (
	"context"

	"vahar.com/go/rest-oati/models"
)

type Repository interface {
	InsertTutorial(cxt context.Context, tutorial *models.Tutorial) error
	GetTutorialById(cxt context.Context, id string) (*models.Tutorial, error)
	GetTutorialByTitle(cxt context.Context, titulo string) (*models.Tutorial, error)
	UpdateTutorial(cxt context.Context, tutorial *models.Tutorial) error
	DeleteTutorial(cxt context.Context, id string) error
	ListTutorial(cxt context.Context, page uint64) ([]*models.Tutorial, error)
	InsertDetail(cxt context.Context, detalle *models.Detalle) error
	GetDetailById(cxt context.Context, id string) (*models.Detalle, error)
	UpdateDetail(cxt context.Context, detalle *models.Detalle) error
	DeleteDetail(cxt context.Context, id string) error
	ListDetail(cxt context.Context, page uint64) ([]*models.Detalle, error)
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

// Buscar por Id
func GetDetailById(cxt context.Context, id string) (*models.Detalle, error) {
	return implementations.GetDetailById(cxt, id)
}

// Modificar tutorial
func UpdateTutorial(cxt context.Context, tutorial *models.Tutorial) error {
	return implementations.UpdateTutorial(cxt, tutorial)
}

// Modificar detalle
func UpdateDetail(cxt context.Context, detalle *models.Detalle) error {
	return implementations.UpdateDetail(cxt, detalle)
}

// Eliminar tutorial
func DeleteTutorial(cxt context.Context, id string) error {
	return implementations.DeleteTutorial(cxt, id)
}

// Eliminar detalle
func DeleteDetail(cxt context.Context, id string) error {
	return implementations.DeleteDetail(cxt, id)
}

// Listar tutorial
func ListTutorial(cxt context.Context, page uint64) ([]*models.Tutorial, error) {
	return implementations.ListTutorial(cxt, page)
}

// Listar detalle
func ListDetail(cxt context.Context, page uint64) ([]*models.Detalle, error) {
	return implementations.ListDetail(cxt, page)
}
func Close() error {
	return implementations.Close()
}
