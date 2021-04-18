package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	issue string
	http.Handler
}

func (s *Server) issueHandler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	length, _ := r.Body.Read(body)
	var jsonBody map[string]interface{}
	json.Unmarshal(body[:length], &jsonBody)
	log.Println("PRINT LOG")
	log.Printf("%v\n", r.Body)
	log.Printf("%v\n", jsonBody)
}

func NewServer() (*Server, error) {
	server := new(Server)

	router := http.NewServeMux()
	router.Handle("/payload", http.HandlerFunc(server.issueHandler))

	server.Handler = router
	return server, nil
}
