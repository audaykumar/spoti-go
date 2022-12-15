package server

import (
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/audaykumar/spoti-go/config"
	"golang.org/x/oauth2"
)

func (s *Server) spotifyAuthenticate(w http.ResponseWriter, req *http.Request) {
	log.Println("spotifyAuthenticate start")
	defer func() {
		log.Println("spotifyAuthenticate Done")
	}()
	state := uuid.New().String()

	w.Header().Add("Location", s.config.OAuth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline))
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (s *Server) handleCallback(w http.ResponseWriter, req *http.Request) {
	log.Println("handleCallback start")
	defer func() {
		log.Println("handleCallback Done")
	}()

	s.oAuthAuthorisationCode = req.URL.Query().Get("code")
	if config.DebugMode {
		log.Println("Received authorisation code:", s.oAuthAuthorisationCode)
	}

	tok, err := s.config.OAuth2Config.Exchange(
		s.context,
		s.oAuthAuthorisationCode,
	)
	if err != nil {
		log.Println("An error occurred while trying to exchange the authorisation code with the Xero API.")
		log.Fatalln(err)
	}
	// Also update the server struct
	s.oAuthToken = tok
	if config.DebugMode {
		log.Println("Got OAuth2 Token from API.")
		log.Println("Access Token:", s.oAuthToken.AccessToken)
		log.Println("Refresh Token:", s.oAuthToken.RefreshToken)
		log.Println("Token expiry:", s.oAuthToken.Expiry.String())
	}

	s.httpClient = *s.config.OAuth2Config.Client(s.context, tok)
	w.Header().Add("Location", "/profile")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
