package config

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	KeyCorrelationId = "correlationId"
	envActiveProfile = "ACTIVE_PROFILE"
)

type ServiceConfig struct {
	LogLevel   string
	ServerPort string
	Kafka      Kafka
}

type Kafka struct {
	Broker string
	Topic  string
}

func NewConfig(ctx context.Context) (ServiceConfig, error) {
	activeProfile := strings.TrimSpace(os.Getenv(envActiveProfile))
	file := "config/default.env"

	if activeProfile != "" {
		file = "config/" + activeProfile + ".env"
		log.Println(ExtractCorrelationId(ctx), "Profile "+activeProfile+" in use")
	} else {
		log.Println(ExtractCorrelationId(ctx), "Profile not set, using default (dev.env)")
	}
	return LoadFileAsConfig(ctx, file)

}

func LoadFileAsConfig(_ context.Context, file string) (ServiceConfig, error) {
	err := godotenv.Load(file)
	if err != nil {
		return ServiceConfig{}, err
	}
	return ServiceConfig{
		LogLevel:   getString("LOG_LEVEL"),
		ServerPort: getString("SERVER_PORT"),
		Kafka: Kafka{
			Topic:  getString("KAFKA_PUBLISH_TOPIC"),
			Broker: getString("KAFKA_BROKER"),
		},
	}, nil
}

func Id(ctx context.Context) string {
	return ExtractCorrelationId(ctx) + " "
}

func ExtractCorrelationId(ctx context.Context) string {
	tokenStr, ok := ctx.Value(KeyCorrelationId).(string)
	if !ok {
		return ""
	}
	return tokenStr
}

func getString(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf(key + " not found in environment")
	return ""
}

func getBool(key string) bool {
	valueString := getString(key)
	if value, err := strconv.ParseBool(valueString); err == nil {
		return value
	}
	log.Fatalf(key + " not found in environment file")
	return false
}
