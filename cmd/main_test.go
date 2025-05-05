package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRootCommandInitialization(t *testing.T) {
    // Test that the rootCmd is properly initialized
    assert.NotNil(t, rootCmd)
    assert.Equal(t, "simple-servicebus-cli", rootCmd.Use)
}

func TestCommandExecution(t *testing.T) {
    // Add tests to verify command execution
    // You might need to capture stdout/stderr for verification
    // For example, you can test the list-queues command
    helpCmd := &cobra.Command{
        Use:   "help",
        Short: "Display help information",
        Run: func(cmd *cobra.Command, args []string) {
            _ = args // Store arguments for potential future use
            err := cmd.Help()
            assert.NoError(t, err, "help command should execute without error")
        },
    }
    err := helpCmd.Execute()
    assert.NoError(t, err, "help command should execute without error")
    

}