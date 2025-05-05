package main

import (
	"fmt"
	"os"

	"simple-servicebus-cli/internal/cli"
)

func main() {
    // Create the root command
    rootCmd := cli.NewRootCommand()

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}