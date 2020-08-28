package api

import (
	"context"
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofrs/uuid"
	"github.com/gorilla/handlers"
	"net/http"
	"strings"
	"timestamp-consumer-service/config"
)

var (
	ErrCorrelationIdNotPresent = errors.New("no " + config.KeyCorrelationId + " header passed")
	ErrCorrelationIdInvalid    = errors.New("invalid " + config.KeyCorrelationId + " header")
)

func (s *server) correlation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationId := r.Header.Get(config.KeyCorrelationId)
		err := validation.Validate(correlationId, validation.Required, is.UUID)
		if err != nil {
			s.respondJSON(w, http.StatusBadRequest, map[string]interface{}{
				"Error": fmt.Sprintf("%s header: %s", config.KeyCorrelationId, err.Error()),
			})
			return
		}

		ctx := context.WithValue(r.Context(), config.KeyCorrelationId, correlationId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *server) validateCorrelationId(correlationId string) error {
	if strings.TrimSpace(correlationId) == "" {
		return ErrCorrelationIdNotPresent
	}
	if _, err := uuid.FromString(correlationId); err != nil {
		return ErrCorrelationIdInvalid
	}
	return nil
}

func (s *server) getCorsConfiguration() (handlers.CORSOption, handlers.CORSOption, handlers.CORSOption) {
	return handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Origin", "Content-Type", "X-Requested-With", "X-Correlation-ID", "X-Client-ID"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost})
}
