package kafka

import (
	"time"

	"github.com/aws/aws-sdk-go/service/kafka"
	"github.com/google/uuid"
)

type Producer struct {
	kafkaClient         *kafka.Kafka
	idGenerator         func() uuid.UUID
	timeNowUTCGenerator func() time.Time
}

type Client struct {
	kafkaClient          *kafka.Kafka
	producerTopic        string
	groupID              string
	consumerTopics       []string
	consumerRetryPeriods []int
}
