package api

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	Error   string `json:"Error"`
	Message string `json:"Message"`
}

func (s *server) respondJSON(writer http.ResponseWriter, code int, payload interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(payload)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(code)
	_, _ = writer.Write(body)
}

func (s *server) respondError(writer http.ResponseWriter, code int, err error, message string) {
	payload := apiError{
		Error:   err.Error(),
		Message: message,
	}
	s.respondJSON(writer, code, payload)
}
