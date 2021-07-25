package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	type Response struct {
		Message string `json:"message"`
	}

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message: "ok",
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, &response)
	})
	r.Get("/header", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, r.Header)
	})

	log.Default().Print("start server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Default().Fatalf("failed to boot server: %s", err.Error())
	}
	log.Default().Printf("exit.")
}
