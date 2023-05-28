package messagequeue

import (
	"context"
	"konfig-go/common/logger"
	"sync"
)

type Message struct {
	Topic string
	Data  interface{}
}

type MessageStore interface {
	SaveMessage(ctx context.Context, topic string, messageData interface{}) error
	GetMessages(ctx context.Context, topic string) ([]Message, error)
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

func (store *InMemoryMessageStore) SaveMessage(ctx context.Context, topic string, messageData interface{}) error {
	store.Lock.Lock()
	defer store.Lock.Unlock()

	message := Message{
		Topic: topic,
		Data:  messageData,
	}

	store.Messages[topic] = append(store.Messages[topic], message)
	logger.Logger.Info(ctx, "Saved message: %s\n", messageData)

	return nil
}

func (store *InMemoryMessageStore) GetMessages(ctx context.Context, topic string) ([]Message, error) {
	store.Lock.RLock()
	defer store.Lock.RUnlock()

	messages := store.Messages[topic]
	return messages, nil
}
