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

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
