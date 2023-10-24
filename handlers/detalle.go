package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	"vahar.com/go/rest-oati/models"
	"vahar.com/go/rest-oati/repository"
	"vahar.com/go/rest-oati/server"
)

type InsertDetailRequest struct {
	TutorialId string `json:"tutorial_id"`
	Autor      string `json:"autor"`
}
type InsertDetailResponse struct {
	Id        string    `json:"id"`
	Autor     string    `json:"autor"`
	CreatedAt time.Time `json:"created_at"`
}

func InsertDetailHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = InsertDetailRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tutorial, err := repository.GetTutorialById(r.Context(), request.TutorialId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var detalle = models.Detalle{
			Id:         id.String(),
			Autor:      request.Autor,
			TutorialId: request.TutorialId,
			CreatedAt:  tutorial.CreatedAt,
		}
		err = repository.InsertDetail(r.Context(), &detalle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(InsertDetailResponse{
			Id:        detalle.Id,
			Autor:     detalle.Autor,
			CreatedAt: detalle.CreatedAt,
		})
	}
}
