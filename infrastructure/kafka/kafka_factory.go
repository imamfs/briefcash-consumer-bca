package kafka

import (
	"context"

	"github.com/IBM/sarama"
)

type KafkaFactory struct {
	Brokers       []string
	GroupID       string
	Topics        []string
	ConsumerGroup sarama.ConsumerGroup
	Handler       sarama.ConsumerGroupHandler
}

func NewKafkaFactory(brokers []string, groupID string, topics []string, handler sarama.ConsumerGroupHandler) (*KafkaFactory, error) {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_8_0_0
	cfg.Consumer.Return.Errors = true
	cfg.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, cfg)
	if err != nil {
		return nil, err
	}

	return &KafkaFactory{
		Brokers:       brokers,
		GroupID:       groupID,
		Topics:        topics,
		ConsumerGroup: consumerGroup,
		Handler:       handler,
	}, nil
}

func (kc *KafkaFactory) Consume(ctx context.Context) error {
	for {
		if err := kc.ConsumerGroup.Consume(ctx, kc.Topics, kc.Handler); err != nil {
			return err
		}
	}
}

func (kc *KafkaFactory) Close() error {
	return kc.ConsumerGroup.Close()
}
