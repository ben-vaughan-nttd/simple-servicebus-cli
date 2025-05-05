package config

import (
    "os"
    "testing"
    
    "github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
    // Save current env var value and restore it after the test
    originalValue := os.Getenv("AZURE_SERVICEBUS_CONNECTION_STRING")
    defer os.Setenv("AZURE_SERVICEBUS_CONNECTION_STRING", originalValue)
    
    // Test when environment variable is not set
    os.Setenv("AZURE_SERVICEBUS_CONNECTION_STRING", "")
    config, err := LoadConfig()
    assert.Error(t, err)
    assert.Nil(t, config)
    
    // Test when environment variable is set
    os.Setenv("AZURE_SERVICEBUS_CONNECTION_STRING", "Endpoint=sb://test.servicebus.windows.net/;SharedAccessKeyName=test;SharedAccessKey=test")
    config, err = LoadConfig()
    assert.NoError(t, err)
    assert.NotNil(t, config)
    assert.Equal(t, "Endpoint=sb://test.servicebus.windows.net/;SharedAccessKeyName=test;SharedAccessKey=test", config.ConnectionString)
}