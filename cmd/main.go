package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MORSE-SERVER: ", log.Ldate|log.Ltime|log.Lshortfile)

	srv := server.NewServer(logger)

	logger.Printf("Starting server on http://localhost:8080")
	if err := srv.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatal("Server failed to start: ", err)
	}
}
