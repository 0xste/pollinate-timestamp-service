package service

import (
	"context"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/Shopify/sarama.v1"
	"time"
	"timestamp-command-service/config"
	"timestamp-command-service/model"
)

type TimestampService interface {
	PublishTimestampRecord(ctx context.Context, timestamp time.Time) (uuid.UUID, error)
}

type timestampService struct {
	log               *logrus.Logger
	kafkaConfig       config.Kafka
	timestampProducer sarama.SyncProducer
}

func NewTimestampService(cfg config.Kafka, logger *logrus.Logger, producer sarama.SyncProducer) *timestampService {
	return &timestampService{
		log:               logger,
		kafkaConfig:       cfg,
		timestampProducer: producer,
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
		EventTimestamp: timestamp,
		CommandId:      commandId.String(),
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

	msg := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(messageKey.String()),
		Value: sarama.StringEncoder(payload),
		Topic: s.kafkaConfig.Topic,
	}
	partition, offset, err := s.timestampProducer.SendMessage(msg)
	if err != nil {
		s.log.Error(err)
		return uuid.Nil, err
	}

	s.log.Infof("message %s published successfully to partition %d with offset %d", messageKey.String(), partition, offset)

	return messageKey, nil
}
