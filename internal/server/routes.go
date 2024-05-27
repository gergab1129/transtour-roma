package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) registerRoutes() http.Handler {
	mux := chi.NewMux()
	mux.Get("/", s.Home)
	mux.Get("/nosotros", s.About)
	mux.Get("/afiliaciones", s.Afiliations)

	mux.Post("/contacto", s.PostMessage)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
