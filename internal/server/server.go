package server

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"transtour-roma/internal/config"
	"transtour-roma/internal/render"
)

type Server struct {
	Port         int
	InProduction bool
	Routes       http.Handler
	Renderer     *render.Renderer
}

func NewServer() *Server {
	appConf := &config.AppConfig{
		InProduction: false,
	}
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
	server.Renderer = render.New(appConf)
	return server
}
