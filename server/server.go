package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/audaykumar/spoti-go/config"
	"golang.org/x/oauth2"
)

const (
	DefaultAppPort = 8888
)

// Server - Represents the running app, and all the bits and pieces we need to make API calls.
type Server struct {
	config                 *config.Config
	context                context.Context
	httpServer             *http.Server
	httpClient             http.Client
	oAuthAuthorisationCode string
	oAuthToken             *oauth2.Token
	userID                 string
}

// New - Returns an instance of the HTTP server.
func New(c *config.Config) *Server {

	c.OAuth2Config.RedirectURL = c.RedirectURI
	if config.DebugMode {
		log.Println("RedirectURL:", c.OAuth2Config.RedirectURL)
	}

	s := &Server{
		config:  c,
		context: context.Background(),
	}

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", s.handleHome)
	r.Get("/login", s.spotifyAuthenticate)
	r.Get("/callback", s.handleCallback)
	r.Get("/profile", s.handleProfile)
	r.Get("/playlists", s.handlePlaylists)

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", DefaultAppPort),
		Handler: r,
	}

	return s
}

// Start - Calls ListenAndServe() on the http server.
func (s *Server) Start() error {
	log.Printf("Hey there! I'm up and running, and can be accessed at: http://localhost:%d\n", DefaultAppPort)
	return s.httpServer.ListenAndServe()
}
