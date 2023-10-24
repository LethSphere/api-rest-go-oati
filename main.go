package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"vahar.com/go/rest-oati/handlers"
	"vahar.com/go/rest-oati/server"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/tutorial", handlers.TutorialRegisterHandler(s)).Methods(http.MethodPost)

	r.HandleFunc("/tutorial", handlers.TutorialConsultHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/detalle", handlers.InsertDetailHandler(s)).Methods(http.MethodPost)
}
