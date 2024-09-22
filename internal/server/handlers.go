package server

import (
	"fmt"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	err := s.Renderer.Render(w, "home.pages.tmpl", s.Renderer.TemplateCache)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error rendering home template. got=%s", err)))
	}
}

func (s *Server) About(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) Afiliations(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) PostMessage(w http.ResponseWriter, r *http.Request) {
}
