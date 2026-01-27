package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestQuoteExists(t *testing.T) {
	quotes := []segment{
		{Text: "Quote one", Attribution: "Author one"},
		{Text: "Quote two", Attribution: "Author two"},
	}

	tests := []struct {
		name     string
		newQuote segment
		want     bool
	}{
		{
			name:     "existing quote",
			newQuote: segment{Text: "Quote one", Attribution: "Author one"},
			want:     true,
		},
		{
			name:     "same text different author",
			newQuote: segment{Text: "Quote one", Attribution: "Different author"},
			want:     true,
		},
		{
			name:     "new quote",
			newQuote: segment{Text: "New quote", Attribution: "New author"},
			want:     false,
		},
		{
			name:     "empty quotes slice",
			newQuote: segment{Text: "Any quote", Attribution: "Any author"},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testQuotes := quotes
			if tt.name == "empty quotes slice" {
				testQuotes = []segment{}
			}
			got := quoteExists(testQuotes, tt.newQuote)
			if got != tt.want {
				t.Errorf("quoteExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadZenlog(t *testing.T) {
	// Save original ZENLOG_FILE
	originalZenlog := ZENLOG_FILE
	defer func() { ZENLOG_FILE = originalZenlog }()

	t.Run("non-existent file returns empty slice", func(t *testing.T) {
		ZENLOG_FILE = "/nonexistent/path/zenlog.json"
		quotes := loadZenlog()
		if len(quotes) != 0 {
			t.Errorf("loadZenlog() returned %d quotes, want 0", len(quotes))
		}
	})

	t.Run("invalid json returns empty slice", func(t *testing.T) {
		tmpDir := t.TempDir()
		ZENLOG_FILE = filepath.Join(tmpDir, "zenlog.json")
		os.WriteFile(ZENLOG_FILE, []byte("invalid json"), 0644)

		quotes := loadZenlog()
		if len(quotes) != 0 {
			t.Errorf("loadZenlog() returned %d quotes, want 0", len(quotes))
		}
	})

	t.Run("valid json returns quotes", func(t *testing.T) {
		tmpDir := t.TempDir()
		ZENLOG_FILE = filepath.Join(tmpDir, "zenlog.json")

		testQuotes := []segment{
			{Text: "Test quote", Attribution: "Test author"},
		}
		data, _ := json.Marshal(testQuotes)
		os.WriteFile(ZENLOG_FILE, data, 0644)

		quotes := loadZenlog()
		if len(quotes) != 1 {
			t.Errorf("loadZenlog() returned %d quotes, want 1", len(quotes))
		}
		if quotes[0].Text != "Test quote" {
			t.Errorf("loadZenlog() text = %q, want %q", quotes[0].Text, "Test quote")
		}
	})
}

func TestAppendToZenlog(t *testing.T) {
	// Save original ZENLOG_FILE
	originalZenlog := ZENLOG_FILE
	defer func() { ZENLOG_FILE = originalZenlog }()

	t.Run("appends to empty file", func(t *testing.T) {
		tmpDir := t.TempDir()
		ZENLOG_FILE = filepath.Join(tmpDir, "zenlog.json")

		quote := segment{Text: "New quote", Attribution: "Author"}
		result := appendToZenlog(quote)

		if !result {
			t.Error("appendToZenlog() returned false, want true")
		}

		quotes := loadZenlog()
		if len(quotes) != 1 {
			t.Errorf("zenlog has %d quotes, want 1", len(quotes))
		}
	})

	t.Run("appends to existing file", func(t *testing.T) {
		tmpDir := t.TempDir()
		ZENLOG_FILE = filepath.Join(tmpDir, "zenlog.json")

		// Create file with one quote
		initial := []segment{{Text: "First quote", Attribution: "Author"}}
		data, _ := json.Marshal(initial)
		os.WriteFile(ZENLOG_FILE, data, 0644)

		// Append new quote
		newQuote := segment{Text: "Second quote", Attribution: "Author"}
		result := appendToZenlog(newQuote)

		if !result {
			t.Error("appendToZenlog() returned false, want true")
		}

		quotes := loadZenlog()
		if len(quotes) != 2 {
			t.Errorf("zenlog has %d quotes, want 2", len(quotes))
		}
	})

	t.Run("rejects duplicate quote", func(t *testing.T) {
		tmpDir := t.TempDir()
		ZENLOG_FILE = filepath.Join(tmpDir, "zenlog.json")

		// Create file with one quote
		initial := []segment{{Text: "Existing quote", Attribution: "Author"}}
		data, _ := json.Marshal(initial)
		os.WriteFile(ZENLOG_FILE, data, 0644)

		// Try to append duplicate
		duplicate := segment{Text: "Existing quote", Attribution: "Different Author"}
		result := appendToZenlog(duplicate)

		if result {
			t.Error("appendToZenlog() returned true for duplicate, want false")
		}

		quotes := loadZenlog()
		if len(quotes) != 1 {
			t.Errorf("zenlog has %d quotes, want 1", len(quotes))
		}
	})
}
