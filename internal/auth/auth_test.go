package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test case 1: API key is present
	t.Run("API key present", func(t *testing.T) {
		header := http.Header{}
		header.Set("Authorization", "ApiKey testkey123")

		apiKey, err := GetAPIKey(header)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if apiKey != "testkey123" {
			t.Errorf("Expected API key 'testkey123', got '%s'", apiKey)
		}
	})

	// Test case 2: API key is missing
	t.Run("API key missing", func(t *testing.T) {
		header := http.Header{}

		_, err := GetAPIKey(header)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})

	// Test case 3: API key is partially missing
	t.Run("API key partially missing", func(t *testing.T) {
		header := http.Header{}
		header.Set("Authorization", "")

		_, err := GetAPIKey(header)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
