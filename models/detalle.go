package models

import "time"

type Detalle struct {
	Id         string    `json:"id"`
	Autor      string    `json:"autor"`
	CreatedAt  time.Time `json:"created_at"`
	TutorialId string    `json:"tutorial_id"`
}
