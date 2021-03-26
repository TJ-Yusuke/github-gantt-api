package main

import (
	server "github-gantt-api/src/server"
	"log"
	"net/http"
)

func main() {
	server, _ := server.NewServer()

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
