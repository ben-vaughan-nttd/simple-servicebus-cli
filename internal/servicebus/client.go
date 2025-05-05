package servicebus

import (
    "context"
    "fmt"

    "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type Client struct {
    // Exported field to allow direct access
    Client *azservicebus.Client
}

func NewClient(connectionString string) (*Client, error) {
    if connectionString == "" {
        return nil, fmt.Errorf("connection string cannot be empty")
    }

    client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create Service Bus client: %v", err)
    }

    return &Client{Client: client}, nil
}

func (c *Client) Close() error {
    return c.Client.Close(context.Background())
}