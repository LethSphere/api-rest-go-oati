package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"vahar.com/go/rest-oati/models"
	"vahar.com/go/rest-oati/repository"
	"vahar.com/go/rest-oati/server"
)

type TutorialUpsertRequest struct {
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
}

type TutorialConsultRequest struct {
	Titulo string `json:"titulo"`
}

type TutorialResponse struct {
	Id          string `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
}
type MessageResponse struct {
	Message string `json:"message"`
}

func TutorialRegisterHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = TutorialUpsertRequest{}
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
			Id:          id.String(),
			Titulo:      request.Titulo,
			Descripcion: request.Descripcion,
			Estado:      request.Estado,
		}
		err = repository.InsertTutorial(r.Context(), &tutorial)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TutorialResponse{
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
		json.NewEncoder(w).Encode(TutorialResponse{
			Id:          tutorial.Id,
			Titulo:      tutorial.Titulo,
			Descripcion: tutorial.Descripcion,
			Estado:      tutorial.Estado,
		})
	}
}

func UpdateTutorialHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var request = TutorialUpsertRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var tutorial = models.Tutorial{
			Id:          params["id"],
			Titulo:      request.Titulo,
			Descripcion: request.Descripcion,
			Estado:      request.Estado,
		}
		err = repository.UpdateTutorial(r.Context(), &tutorial)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TutorialResponse{
			Id:          tutorial.Id,
			Titulo:      tutorial.Titulo,
			Descripcion: tutorial.Descripcion,
			Estado:      tutorial.Estado,
		})
	}
}

func DeleteTutorialHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		err := repository.DeleteTutorial(r.Context(), params["id"])
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

func ListTutorialHandler(s server.Server) http.HandlerFunc {
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
		tutorials, err := repository.ListTutorial(r.Context(), page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tutorials)
	}
}
