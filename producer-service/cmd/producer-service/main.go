package main

import (
	"log"
	"os"

	"github.com/suhail34/pdf-compressor/producer-service/internal/server"
)

func main() {
	srv := server.NewServer()
	if err := srv.Run(); err != nil {
		log.Print("Error starting the server : ", err)
		os.Exit(1)
	}
}
