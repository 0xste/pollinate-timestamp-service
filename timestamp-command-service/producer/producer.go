package producer

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaProducer interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
}

type producer struct {
	kafkaProducer KafkaProducer
	logger        logrus.Logger
}

func NewProducer(cfg ClusterConfig, logger logrus.Logger) (*producer, error) {
	kafkaProducerConfig := kafka.ConfigMap{
		"bootstrap.servers":                   cfg.Brokers,
		"enable.ssl.certificate.verification": cfg.SSLCertificateVerification,
		"enable.idempotence":                  true,
	}
	if cfg.SSLCertificateAuthorityLocation != "" {
		kafkaProducerConfig["ssl.ca.location"] = cfg.SSLCertificateAuthorityLocation
	}
	if cfg.SASLMechanism != "" {
		kafkaProducerConfig["sasl.mechanism"] = cfg.SASLMechanism
	}
	if cfg.SecurityProtocol != "" {
		kafkaProducerConfig["security.protocol"] = cfg.SecurityProtocol
	}
	if cfg.Username != "" {
		kafkaProducerConfig["sasl.username"] = cfg.Username
	}
	if cfg.Password != "" {
		kafkaProducerConfig["sasl.password"] = cfg.Password
	}

	kafkaProducer, err := kafka.NewProducer(&kafkaProducerConfig)
	if err != nil {
		return &producer{}, err
	}

	return &producer{kafkaProducer: kafkaProducer, logger: logger}, nil
}

func (p *producer) publish(topic string, key, value []byte, headers []kafka.Header) error {
	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:   value,
		Key:     key,
		Headers: headers,
	}

	deliveryChan := make(chan kafka.Event, 1)
	err := p.kafkaProducer.Produce(&msg, deliveryChan)
	if err != nil {
		return err
	}

	// block until we get a delivery event to make synchronous
	event := <-deliveryChan
	delivery := event.(*kafka.Message)

	if delivery.TopicPartition.Error != nil {
		return err
	}

	return nil
}
