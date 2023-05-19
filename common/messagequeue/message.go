package messagequeue

import (
	"fmt"
	"sync"
)

type Message struct {
	Topic string
	Data  string
}

type MessageStore interface {
	SaveMessage(topic, messageData string) error
	GetMessages(topic string) ([]Message, error)
}

type InMemoryMessageStore struct {
	Messages map[string][]Message
	Lock     sync.RWMutex
}

func NewInMemoryMessageStore() *InMemoryMessageStore {
	return &InMemoryMessageStore{
		Messages: make(map[string][]Message),
	}
}

func (store *InMemoryMessageStore) SaveMessage(topic, messageData string) error {
	store.Lock.Lock()
	defer store.Lock.Unlock()

	message := Message{
		Topic: topic,
		Data:  messageData,
	}

	store.Messages[topic] = append(store.Messages[topic], message)
	fmt.Sprintf("Saved message: %s\n", messageData)
	return nil
}

func (store *InMemoryMessageStore) GetMessages(topic string) ([]Message, error) {
	store.Lock.RLock()
	defer store.Lock.RUnlock()

	messages := store.Messages[topic]
	return messages, nil
}
