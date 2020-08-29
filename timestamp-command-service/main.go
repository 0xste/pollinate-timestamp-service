package main

import (
	"context"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"timestamp-command-service/api"
	"timestamp-command-service/config"
)

func main() {
	globalContext := context.WithValue(context.Background(), config.KeyCorrelationId, uuid.Must(uuid.NewV4()).String())

	if err := run(globalContext); err != nil {
		log.Fatal(err.Error())
	}
}

func run(ctx context.Context) error {
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}
	log.Println(cfg)

	server := api.NewServer(ctx, cfg)
	return http.ListenAndServe(":"+cfg.ServerPort, server)
}
