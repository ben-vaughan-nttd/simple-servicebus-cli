package servicebus

import (
    "context"
    "fmt"
    
    "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
    "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/admin"
)

// QueueManager manages operations related to Service Bus queues.
type QueueManager struct {
    client      *azservicebus.Client
    adminClient *admin.Client
}

// NewQueueManager creates a new QueueManager instance.
func NewQueueManager(connectionString string) (*QueueManager, error) {
    client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create service bus client: %v", err)
    }
    
    adminClient, err := admin.NewClientFromConnectionString(connectionString, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create admin client: %v", err)
    }
    
    return &QueueManager{
        client:      client,
        adminClient: adminClient,
    }, nil
}

// ListQueues enumerates all queues in the Service Bus namespace.
func (qm *QueueManager) ListQueues(ctx context.Context) ([]string, error) {
    // Use the admin client to list queues
    pager := qm.adminClient.NewListQueuesRuntimePropertiesPager(nil)
    
    var queueNames []string
    for pager.More() {
        page, err := pager.NextPage(ctx)
        if err != nil {
            return nil, fmt.Errorf("failed to list queues: %v", err)
        }
        
        for _, queue := range page.QueueRuntimeProperties {
            // Fix: Use the correct field for queue name based on Azure SDK
            queueNames = append(queueNames, queue.QueueName) // Updated from *queue.Name
        }
    }
    
    return queueNames, nil
}

// GetReceiver creates a receiver for the specified queue.
func (qm *QueueManager) GetReceiver(ctx context.Context, queueName string) (*azservicebus.Receiver, error) {
    receiver, err := qm.client.NewReceiverForQueue(queueName, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create receiver for queue %s: %v", queueName, err)
    }
    
    return receiver, nil
}