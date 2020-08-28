package service

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type TimestampService interface {
	PublishTimestampRecord(ctx context.Context, timestamp time.Time) (uuid.UUID, error)
}

type timestampService struct {
	log               *logrus.Logger
}

func NewTimestampService(logger *logrus.Logger) *timestampService {
	return &timestampService{
		log:         logger,
	}
}

func (s *timestampService) PublishTimestampRecord(ctx context.Context, timestamp time.Time) (uuid.UUID, error) {
	return uuid.Nil, nil
}
