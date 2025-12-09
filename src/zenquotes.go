package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// ZenQuote represents a single quote from the ZenQuotes API
type ZenQuote struct {
	Text   string `json:"q"`
	Author string `json:"a"`
}

// quoteCache holds the last N quotes fetched from the API
var quoteCache []segment
var cacheMaxSize = 10

// fallbackQuote is used when API fails and cache is empty
var fallbackQuote = segment{
	Text:        "The only way to do great work is to love what you do.",
	Attribution: "Steve Jobs",
}

// fetchZenQuote fetches a random quote from the ZenQuotes API
func fetchZenQuote() (segment, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://zenquotes.io/api/random")
	if err != nil {
		return segment{}, err
	}
	defer resp.Body.Close()

	var quotes []ZenQuote
	if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		return segment{}, err
	}

	if len(quotes) == 0 {
		return segment{}, err
	}

	// Convert to segment format
	result := segment{
		Text:        quotes[0].Text,
		Attribution: quotes[0].Author,
	}

	// Add to cache
	quoteCache = append(quoteCache, result)
	if len(quoteCache) > cacheMaxSize {
		quoteCache = quoteCache[1:]
	}

	return result, nil
}

// getQuoteWithFallback attempts to fetch from API, falls back to cache,
// then falls back to default quote
func getQuoteWithFallback() segment {
	// Try to fetch from API
	quote, err := fetchZenQuote()
	if err == nil {
		return quote
	}

	// Fall back to cache
	if len(quoteCache) > 0 {
		idx := rand.Intn(len(quoteCache))
		return quoteCache[idx]
	}

	// Fall back to hardcoded quote
	return fallbackQuote
}

// generateZenQuotesTest returns a function that generates a new quote
// segment each time it's called
func generateZenQuotesTest() func() []segment {
	return func() []segment {
		return []segment{getQuoteWithFallback()}
	}
}
