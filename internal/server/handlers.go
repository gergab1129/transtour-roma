package server

import "net/http"

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func (s *Server) About(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) Afiliations(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) PostMessage(w http.ResponseWriter, r *http.Request) {

}
