package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"vahar.com/go/rest-oati/models"
	"vahar.com/go/rest-oati/repository"
	"vahar.com/go/rest-oati/server"
)

type UpsertDetailRequest struct {
	TutorialId string `json:"tutorial_id"`
	Autor      string `json:"autor"`
}
type DetailResponse struct {
	Id        string    `json:"id"`
	Autor     string    `json:"autor"`
	CreatedAt time.Time `json:"created_at"`
}

func InsertDetailHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = UpsertDetailRequest{}
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
		detalle := models.Detalle{
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
		json.NewEncoder(w).Encode(DetailResponse{
			Id:        detalle.Id,
			Autor:     detalle.Autor,
			CreatedAt: detalle.CreatedAt,
		})
	}
}

func GetDetailByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		detalle, err := repository.GetDetailById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(detalle)
	}
}

func UpdateDetailHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var request = UpsertDetailRequest{}
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

		detalle := models.Detalle{
			Id:         params["id"],
			Autor:      request.Autor,
			TutorialId: request.TutorialId,
			CreatedAt:  tutorial.CreatedAt,
		}
		err = repository.UpdateDetail(r.Context(), &detalle)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(DetailResponse{
			Id:        detalle.Id,
			Autor:     detalle.Autor,
			CreatedAt: detalle.CreatedAt,
		})
	}
}
func DeleteDetailHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		err := repository.DeleteDetail(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(MessageResponse{
			Message: "Deleted tutorial",
		})
	}
}

func ListDetailHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		pageStr := r.URL.Query().Get("page")
		var page = uint64(0)
		if pageStr != "" {
			page, err = strconv.ParseUint(pageStr, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		detalles, err := repository.ListDetail(r.Context(), page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(detalles)
	}
}
