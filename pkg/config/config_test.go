package config

import (
	"os"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	os.Setenv("DEBUG", "true")
	os.Setenv("METRICS_PORT", "9000")
	os.Setenv("SERVER_PORT", "3000")

	LoadConfiguration()

	if !CFG.Debug {
		t.Errorf("Expected Debug to be true, got %v", CFG.Debug)
	}
	if CFG.MetricsPort != 9000 {
		t.Errorf("Expected MetricsPort to be 9000, got %d", CFG.MetricsPort)
	}
	if CFG.ServerPort != 3000 {
		t.Errorf("Expected ServerPort to be 3000, got %d", CFG.ServerPort)
	}
}

func TestGetEnvOrDefault(t *testing.T) {
	key := "TEST_KEY"
	defaultValue := "default"

	value := getEnvOrDefault(key, defaultValue)
	if value != defaultValue {
		t.Errorf("Expected '%s', got '%s'", defaultValue, value)
	}

	expectedValue := "test-value"
	os.Setenv(key, expectedValue)
	value = getEnvOrDefault(key, defaultValue)
	if value != expectedValue {
		t.Errorf("Expected '%s', got '%s'", expectedValue, value)
	}
}
