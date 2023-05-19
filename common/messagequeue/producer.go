package messagequeue

import "fmt"

type IProducer interface {
	Produce(topic, messageData string) error
}

type Producer struct {
	MessageStore MessageStore
}

func NewProducer(messageStore MessageStore) *Producer {
	return &Producer{
		MessageStore: messageStore,
	}
}
func (producer *Producer) Produce(topic, messageData string) error {
	err := producer.MessageStore.SaveMessage(topic, messageData)
	if err != nil {
		fmt.Sprintf("Failed to save message: %v\n", err)
	}
	return err
}
