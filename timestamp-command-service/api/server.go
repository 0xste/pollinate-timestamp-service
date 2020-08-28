package api

import (
	"context"
	"crypto/tls"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
	"timestamp-command-service/config"
	"timestamp-command-service/service"
)

const defaultLogLevel = logrus.InfoLevel

type server struct {
	router           *mux.Router
	log              *logrus.Logger
	timestampService service.TimestampService
}

func NewServer(ctx context.Context, cfg config.ServiceConfig) *server {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       &tls.Config{},
	}

	logLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logLevel = defaultLogLevel
	}
	logger := NewLogrus(logLevel)
	s := &server{
		router:           mux.NewRouter(),
		log:              logger,
		timestampService: service.NewTimestampService(cfg.Kafka, logger, dialer),
	}
	s.routes()
	logger.Info(config.Id(ctx), "Starting HTTP server on :"+cfg.ServerPort)
	return s
}

func NewLogrus(level logrus.Level) *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stderr,
		Level:     level,
		Formatter: &logrus.JSONFormatter{},
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Authorization, X-Correlation-ID, X-Client-ID")
	// Stop here if its Preflighted OPTIONS request. Otherwise CORS will reject it
	if r.Method == "OPTIONS" {
		return
	}
	s.router.ServeHTTP(w, r)
}