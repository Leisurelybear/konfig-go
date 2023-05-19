package messagequeue

import (
	"fmt"
	"time"
)

type IConsumer interface {
	Consume()
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

func (consumer *Consumer) Consume() {
	for {
		messages, err := consumer.MessageStore.GetMessages(consumer.Topic)
		if err != nil {
			fmt.Sprintf("Failed to get messages: %v\n", err)
			return
		}

		for i := consumer.Offset; i < len(messages); i++ {
			message := messages[i]
			fmt.Sprintf("Consumer %d - Consumed: %s\n", consumer.ID, message.Data)
			consumer.Offset = i + 1

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
