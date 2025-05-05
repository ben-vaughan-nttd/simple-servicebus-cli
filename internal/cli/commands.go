package cli

import (
	"context"
	"fmt"

	"simple-servicebus-cli/internal/config"
	"simple-servicebus-cli/internal/servicebus"

	"github.com/spf13/cobra"
)

var (
    queueName string
    maxMessages int32 = 10 // Default value for maximum messages to peek
)

func NewRootCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "simple-servicebus-cli",
        Short: "A CLI tool for interacting with Azure Service Bus",
        Long:  `This tool allows you to enumerate service bus queues and peek at messages in those queues.`,
    }

    cmd.AddCommand(newListQueuesCommand())
    cmd.AddCommand(newPeekMessageCommand())

    return cmd
}

func newListQueuesCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "list-queues",
        Short: "List all Service Bus queues",
        RunE:  listQueues,
    }
}

func listQueues(_ *cobra.Command, args []string) error {
    // Get configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        return fmt.Errorf("failed to load configuration: %w", err)
    }

    // Store arguments for potential future use
    _ = args

    // Create queue manager
    queueManager, err := servicebus.NewQueueManager(cfg.ConnectionString)
    if err != nil {
        return fmt.Errorf("failed to create queue manager: %w", err)
    }

    // List queues
    ctx := context.Background()
    queues, err := queueManager.ListQueues(ctx)
    if err != nil {
        return fmt.Errorf("failed to list queues: %w", err)
    }

    // Print output
    if len(queues) == 0 {
        fmt.Println("No queues found")
    } else {
        fmt.Println("Available queues:")
        for _, queue := range queues {
            fmt.Printf("- %s\n", queue)
        }
    }

    return nil
}

func newPeekMessageCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "peek-message",
        Short: "Peek at messages in a Service Bus queue",
        RunE:  peekMessage,
    }

    cmd.Flags().StringVarP(&queueName, "queue", "q", "", "Name of the queue to peek messages from")
    cmd.Flags().Int32VarP(&maxMessages, "max", "m", 10, "Maximum number of messages to peek")
    if err := cmd.MarkFlagRequired("queue"); err != nil {
        // Since this is during initialization, panic is appropriate
        panic(fmt.Sprintf("failed to mark 'queue' flag as required: %v", err))
    }
    return cmd
}

func peekMessage(_ *cobra.Command, args []string) error {
    // Get configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        return fmt.Errorf("failed to load configuration: %w", err)
    }

    // Store arguments for potential future use
    _ = args

    // Create client
    client, err := servicebus.NewClient(cfg.ConnectionString)
    if err != nil {
        return fmt.Errorf("failed to create Service Bus client: %w", err)
    }
    defer client.Close()

    // Create message service
    messageService := servicebus.NewMessageService(client.Client, queueName)

    // Peek messages
    ctx := context.Background()
    messages, err := messageService.PeekMessages(ctx, maxMessages)
    if err != nil {
        return fmt.Errorf("failed to peek messages: %w", err)
    }

    // Print output
    if len(messages) == 0 {
        fmt.Println("No messages found in queue:", queueName)
    } else {
        fmt.Printf("Messages in queue %s:\n", queueName)
        for i, message := range messages {
            fmt.Printf("Message #%d:\n", i+1)
            fmt.Printf("  ID: %s\n", message.MessageID)
            fmt.Printf("  Enqueued: %s\n", message.EnqueuedTime.Format("2006-01-02 15:04:05"))
            fmt.Printf("  Body: %s\n", string(message.Body))
            fmt.Println()
        }
    }

    return nil
}