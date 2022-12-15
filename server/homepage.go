package server

import (
	"log"
	"net/http"
)

func (s *Server) handleHome(w http.ResponseWriter, req *http.Request) {
	log.Println("handleHome start")
	defer func() {
		log.Println("handleHome Done")
	}()
}
