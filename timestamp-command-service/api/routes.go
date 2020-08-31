package api

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	pathBase              string = "/api"
	pathVersion1          string = "/v1"
	pathPostTimestamp     string = "/timestamp/{timestamp}"
	pathPostTimestampSpec string = "/app"
	pathHealth string = "/health"
)

func (s *server) routes() {
	s.apiRoutes()
	s.apiMiddleware(s.router)
}

func (s *server) apiRoutes() {
	s.router.HandleFunc(pathPostTimestampSpec, s.submitTimestampRecord).Methods(http.MethodPost)
	s.router.HandleFunc(pathHealth, s.health).Methods(http.MethodGet)
	sub := s.router.PathPrefix(pathBase + pathVersion1).Subrouter()
	sub.HandleFunc(pathPostTimestamp, s.submitTimestampRecord).Methods(http.MethodPost)
}

func (s *server) apiMiddleware(r *mux.Router) {
	r.Use(handlers.ProxyHeaders)
	r.Use(s.correlation)
}
