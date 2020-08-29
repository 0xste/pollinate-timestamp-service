package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"timestamp-command-service/config"
)

// submitTimestampRecord godoc
// @Summary Publishes a timestamp record to kafka, the command segment from CQRS design.
// @Tags timestamp-command-service
// @Description
// @Produce json
// @Param X-Correlation-ID header string true "X-Correlation-ID" minlength(36) maxlength(36) default(UUIDv4 -> "c7a84eb2-40f7-4806-b719-c9655e31ad2f")
// @Param timestamp query string true "timestamp" default(RFC3339 -> "2020-08-28T08:40:49.546300996Z")
// @Success 200 {object} idResponse
// @Failure 400 {object} apiError
// @Failure 500 {object} apiError
// @Router /api/v1/timestamp/{timestamp} [get]
func (s *server) submitTimestampRecord(w http.ResponseWriter, r *http.Request) {
	s.log.Debug(config.Id(r.Context()), "Enter: submitTimestampRecord")
	defer s.log.Debug(config.Id(r.Context()), "Exit: submitTimestampRecord")

	ts, err := time.Parse(time.RFC3339, mux.Vars(r)["timestamp"])
	if err != nil {
		s.log.Error(err.Error())
		s.respondError(w, http.StatusBadRequest, err, fmt.Sprintf("invalid timestamp passed, please use %s", time.RFC3339))
		return
	}

	if record, err := s.timestampService.PublishTimestampRecord(r.Context(), ts); err != nil {
		s.log.Error(err.Error())
		s.respondError(w, http.StatusInternalServerError, err, fmt.Sprintf("failed to publish message"))
		return
	} else {
		s.log.Infof("record published successfully for id: %s", record.String())
		s.respondJSON(w, http.StatusOK, idResponse{
			Id: record.String(),
		})
	}
}

type idResponse struct {
	Id string `json:"id"`
}
