package models

import "time"

type Tutorial struct {
	Id          string    `json:"id"`
	Titulo      string    `json:"titulo"`
	Descripcion string    `json:"descripcion"`
	Estado      string    `json:"estado"`
	CreatedAt   time.Time `json:"created_at"`
}
