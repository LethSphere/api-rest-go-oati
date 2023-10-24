package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/segmentio/ksuid"
	"vahar.com/go/rest-oati/models"
	"vahar.com/go/rest-oati/repository"
	"vahar.com/go/rest-oati/server"
)

type TutorialRegisterRequest struct {
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
}

type TutorialRegisterResponse struct {
	Id          string `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
}
type TutorialConsultRequest struct {
	Titulo string `json:"titulo"`
}

type TutorialConsultResponse struct {
	Id          string `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
}

func TutorialRegisterHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = TutorialRegisterRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var tutorial = models.Tutorial{
			Titulo:      request.Titulo,
			Descripcion: request.Descripcion,
			Estado:      request.Estado,
			Id:          id.String(),
		}
		err = repository.InsertTutorial(r.Context(), &tutorial)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TutorialRegisterResponse{
			Id:          tutorial.Id,
			Titulo:      tutorial.Titulo,
			Descripcion: tutorial.Descripcion,
			Estado:      tutorial.Estado,
		})
	}
}

func TutorialConsultHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = TutorialConsultRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tutorial, err := repository.GetTutorialByTitle(r.Context(), request.Titulo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if tutorial == nil {
			http.Error(w, "Not fund", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TutorialConsultResponse{
			Id:          tutorial.Id,
			Titulo:      tutorial.Titulo,
			Descripcion: tutorial.Descripcion,
			Estado:      tutorial.Estado,
		})
	}
}
