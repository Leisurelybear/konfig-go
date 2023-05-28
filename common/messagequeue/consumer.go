package messagequeue

import (
	"context"
	"fmt"
	"konfig-go/common/logger"
	"time"
)

type IConsumer interface {
	Consume(ctx context.Context, handleFunc func(data interface{}) error)
	Close()
}

type Consumer struct {
	ID           int
	Topic        string
	Offset       int
	MessageStore MessageStore
	doneCh       chan bool
}

func NewConsumer(id int, topic string, offset int, messageStore MessageStore) *Consumer {
	return &Consumer{
		ID:           id,
		Topic:        topic,
		Offset:       offset,
		MessageStore: messageStore,
		doneCh:       make(chan bool),
	}
}

func (consumer *Consumer) Consume(ctx context.Context, handleFunc func(data interface{}) error) {
	for {
		messages, err := consumer.MessageStore.GetMessages(ctx, consumer.Topic)
		if err != nil {
			fmt.Sprintf("Failed to get messages: %v\n", err)
			return
		}

		for i := consumer.Offset; i < len(messages); i++ {
			message := messages[i]
			er := handleFunc(message.Data)
			if er != nil {
				logger.Logger.Error(ctx, "error to consume %d - err: %v\n", consumer.ID, er)
				break
			} else {
				consumer.Offset = i + 1
				logger.Logger.Info(ctx, "Consumer %d - Consumed: %s\n", consumer.ID, message.Data)
			}

			select {
			case <-consumer.doneCh:
				fmt.Sprintf("Consumer %d - Done\n", consumer.ID)
				return
			default:
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func (consumer *Consumer) Close() {
	close(consumer.doneCh)
}
