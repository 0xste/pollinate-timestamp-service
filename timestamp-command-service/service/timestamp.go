package service

import (
	"context"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"time"
	"timestamp-command-service/config"
	"timestamp-command-service/model"
)

type TimestampService interface {
	PublishTimestampRecord(ctx context.Context, timestamp time.Time) (uuid.UUID, error)
}

type timestampService struct {
	log         *logrus.Logger
	kafkaConfig config.Kafka
	timestampProducer *kafka.Writer
}

func NewTimestampService(cfg config.Kafka, logger *logrus.Logger) *timestampService {
	return &timestampService{
		log:         logger,
		kafkaConfig: cfg,
		timestampProducer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{cfg.Broker},
			Topic:   cfg.Topic,
			Balancer: &kafka.CRC32Balancer{},
		}),
	}
}

func (s *timestampService) PublishTimestampRecord(ctx context.Context, timestamp time.Time) (uuid.UUID, error) {
	s.log.Debug(config.Id(ctx), "Enter: PublishTimestampRecord")
	defer s.log.Debug(config.Id(ctx), "Exit: PublishTimestampRecord")

	commandId, err := uuid.NewV4()
	if err != nil {
		s.log.Error(err)
		return uuid.Nil, err
	}
	payload, err := json.Marshal(model.Timestamp{
		EventTimestamp:   timestamp,
		CommandTimestamp: time.Now(),
		CommandId:        commandId.String(),
	})
	if err != nil {
		s.log.Error(err)
		return uuid.Nil, err
	}

	messageKey, err := uuid.NewV4()
	if err != nil {
		s.log.Error(err)
		return uuid.Nil, err
	}

	if err := s.timestampProducer.WriteMessages(ctx, kafka.Message{
		Key:   messageKey.Bytes(),
		Value: payload,
	}) ; err != nil{
		s.log.Error(err)
		return uuid.Nil, err
	}

	return messageKey, nil
}

