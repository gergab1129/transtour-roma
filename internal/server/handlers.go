package server

import "net/http"

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
