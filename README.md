# Azure Service Bus CLI Tool

This project is a command-line interface (CLI) tool for interacting with Azure Service Bus. It allows users to enumerate service bus queues and peek at messages within those queues.

## Features

- List all service bus queues
- Peek at messages in a specified queue

## Prerequisites

- Go 1.16 or later
- An Azure account with access to Azure Service Bus

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/ben-vaughan-ntt/simple-servicebus-cli.git
   ```

2. Navigate to the project directory:

   ```shell
   cd simple-servicebus-cli
   ```

3. Install the dependencies:

   ```shell
   go mod tidy
   ```

## Usage

To run the CLI tool, use the following command:

```shell
go run cmd/main.go [command]
```

You must have the service bus connection string in an environment variable called `AZURE_SERVICEBUS_CONNECTION_STRING`.

### Commands

- `list`: Enumerates all service bus queues.
- `peek <queue-name>`: Peeks at messages in the specified queue.

## Configuration

The tool requires configuration settings to connect to Azure Service Bus. You can set these in the `config.go` file or through environment variables.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
