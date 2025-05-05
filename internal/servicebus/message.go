package servicebus

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

// Message represents a Service Bus message with relevant properties.
type Message struct {
	MessageID            string
	Body                 []byte
	ApplicationProperties map[string]interface{}
	EnqueuedTime         time.Time
}

// MessageService provides methods to interact with messages in a Service Bus queue.
type MessageService struct {
	client *azservicebus.Client
	queue  string
}

// NewMessageService creates a new instance of MessageService.
func NewMessageService(client *azservicebus.Client, queue string) *MessageService {
	return &MessageService{
		client: client,
		queue:  queue,
	}
}

// PeekMessages retrieves a specified number of messages from the queue without removing them.
func (ms *MessageService) PeekMessages(ctx context.Context, maxMessages int32) ([]Message, error) {
	receiver, err := ms.client.NewReceiverForQueue(ms.queue, nil)
	if err != nil {
		return nil, err
	}
	defer receiver.Close(ctx)

	messages, err := receiver.PeekMessages(ctx, int(maxMessages), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to peek messages: %v", err)
	}

	result := make([]Message, 0, len(messages))
	for _, msg := range messages {
		message := Message{
			MessageID:            msg.MessageID,
			Body:                 msg.Body,
			ApplicationProperties: msg.ApplicationProperties,
			EnqueuedTime:         time.Now(),
		}
		if msg.EnqueuedTime != nil {
			message.EnqueuedTime = *msg.EnqueuedTime
		}
		result = append(result, message)
	}

	return result, nil
}