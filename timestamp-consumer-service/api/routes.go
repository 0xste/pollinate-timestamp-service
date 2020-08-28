package api

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	pathBase          string = "/api"
	pathVersion1      string = "/v1"
	pathPostTimestamp string = "/timestamp/{timestamp}"
)

func (s *server) routes() {
	s.apiMiddleware(s.router)
	s.apiRoutes()
}

func (s *server) apiRoutes() {
	sub := s.router.PathPrefix(pathBase + pathVersion1).Subrouter()
	sub.HandleFunc(pathPostTimestamp, s.submitTimestampRecord).Methods(http.MethodPost)
}

func (s *server) apiMiddleware(r *mux.Router) {
	r.Use(handlers.ProxyHeaders)
	r.Use(s.correlation)
}
