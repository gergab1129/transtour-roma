package main

import (
	"fmt"
	"log"
	"net/http"
	"transtour-roma/internal/server"
)

func main() {

	server := server.NewServer()
	fmt.Printf("Listening on http://localhost:%d\n", server.Port)
	ser := &http.Server{
		Addr:    fmt.Sprintf(":%d", server.Port),
		Handler: server.Routes,
	}

	err := ser.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed server initialization. error=%s", err)
	}
}
