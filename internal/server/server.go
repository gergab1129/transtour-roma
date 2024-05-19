package server

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	Port         int
	InProduction bool
	Routes       http.Handler
}

func NewServer() *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("failed to convert port to int error. got=%s", err)
	}
	server := &Server{
		Port: serverPort,
	}
	server.Routes = server.registerRoutes()
	return server
}
