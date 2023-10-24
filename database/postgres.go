package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"vahar.com/go/rest-oati/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) InsertTutorial(cxt context.Context,
	tutorial *models.Tutorial) error {
	_, err := repo.db.ExecContext(cxt,
		"INSERT INTO tutorials (id, titulo, descripcion, estado) VALUES ($1, $2, $3, $4)",
		tutorial.Id, tutorial.Titulo, tutorial.Descripcion, tutorial.Estado)
	return err
}

func (repo *PostgresRepository) GetTutorialById(cxt context.Context,
	id string) (*models.Tutorial, error) {
	rows, err := repo.db.QueryContext(cxt,
		"SELECT id, titulo, descripcion, estado, created_at FROM tutorials WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var tutorial = models.Tutorial{}
	for rows.Next() {
		if err := rows.Scan(&tutorial.Id, &tutorial.Titulo,
			&tutorial.Descripcion, &tutorial.Estado); err == nil {
			return &tutorial, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &tutorial, nil
}

func (repo *PostgresRepository) GetTutorialByTitle(cxt context.Context,
	titulo string) (*models.Tutorial, error) {
	rows, err := repo.db.QueryContext(cxt,
		"SELECT id, titulo, descripcion, estado FROM tutorials WHERE titulo = $1", titulo)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var tutorial = models.Tutorial{}
	for rows.Next() {
		if err := rows.Scan(&tutorial.Id, &tutorial.Titulo,
			&tutorial.Descripcion, &tutorial.Estado); err == nil {
			return &tutorial, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &tutorial, nil
}

func (repo *PostgresRepository) InsertDetail(cxt context.Context,
	detalle *models.Detalle) error {
	_, err := repo.db.ExecContext(cxt,
		"INSERT INTO detalles (id, autor, created_at, tutorials_id) VALUES ($1, $2, $3, $4)",
		detalle.Id, detalle.Autor, detalle.CreatedAt, detalle.TutorialId)
	return err
}

func (repo *PostgresRepository) GetDetailById(cxt context.Context,
	id string) (*models.Detalle, error) {
	rows, err := repo.db.QueryContext(cxt,
		"SELECT id, autor, created_at, tutorial_id FROM detalles WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var detalle = models.Detalle{}
	for rows.Next() {
		if err := rows.Scan(&detalle.Id, &detalle.Autor, &detalle.CreatedAt,
			&detalle.TutorialId); err == nil {
			return &detalle, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &detalle, nil
}

func (repo *PostgresRepository) UpdateTutorial(cxt context.Context,
	tutorial *models.Tutorial) error {
	_, err := repo.db.ExecContext(cxt, "UPDATE tutorials SET titulo = $1, descripcion = $2 WHERE id = $3",
		tutorial.Titulo, tutorial.Descripcion, tutorial.Id)
	return err
}

func (repo *PostgresRepository) UpdateDetail(cxt context.Context,
	detalle *models.Detalle) error {
	_, err := repo.db.ExecContext(cxt, "UPDATE detalles SET autor = $1 WHERE id = $2",
		detalle.Autor, detalle.Id)
	return err
}
func (repo *PostgresRepository) DeleteTutorial(cxt context.Context, id string) error {
	_, err := repo.db.ExecContext(cxt, "DELETE FROM tutorials WHERE id = $1",
		id)
	return err
}
func (repo *PostgresRepository) DeleteDetail(cxt context.Context, id string) error {
	_, err := repo.db.ExecContext(cxt, "DELETE FROM detalles WHERE id = $1",
		id)
	return err
}
func (repo *PostgresRepository) ListTutorial(cxt context.Context, page uint64) ([]*models.Tutorial, error) {
	rows, err := repo.db.QueryContext(cxt, "SELECT id, titulo, descripcion, estado FROM tutorials LIMIT $1 OFFSET $2", 10, page*10)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var tutorials []*models.Tutorial
	for rows.Next() {
		var tutorial = models.Tutorial{}
		if err = rows.Scan(&tutorial.Id, &tutorial.Titulo, &tutorial.Descripcion,
			&tutorial.Estado); err == nil {
			tutorials = append(tutorials, &tutorial)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tutorials, nil
}
func (repo *PostgresRepository) ListDetail(cxt context.Context, page uint64) ([]*models.Detalle, error) {
	rows, err := repo.db.QueryContext(cxt, "SELECT id, autor, created_at FROM detalles LIMIT $1 OFFSET $2", 10, page*10)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var detalles []*models.Detalle
	for rows.Next() {
		var detalle = models.Detalle{}
		if err = rows.Scan(&detalle.Id, &detalle.Autor, &detalle.CreatedAt); err == nil {
			detalles = append(detalles, &detalle)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return detalles, nil
}
func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
