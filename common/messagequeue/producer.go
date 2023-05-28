package messagequeue

import (
	"context"
	"konfig-go/common/logger"
)

type IProducer interface {
	Produce(ctx context.Context, topic string, messageData interface{}) error
}

type Producer struct {
	MessageStore MessageStore
}

func NewProducer(messageStore MessageStore) *Producer {
	return &Producer{
		MessageStore: messageStore,
	}
}
func (producer *Producer) Produce(ctx context.Context, topic string, messageData interface{}) error {
	err := producer.MessageStore.SaveMessage(topic, messageData)
	if err != nil {
		logger.Logger.Error(ctx, "Failed to save message: %v\n", err)
	}
	return err
}
