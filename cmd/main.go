package cmd

import (
	"simple-servicebus-cli/internal/cli"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "simple-servicebus-cli",
    Short: "A CLI tool for Azure Service Bus",
    // Add more command configuration here
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    rootCmd = cli.NewRootCommand()
    // Add subcommands and flags here
}