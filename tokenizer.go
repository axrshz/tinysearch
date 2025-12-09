package main

import (
	"strings"
	"unicode"
)

// Tokenizer is a tool that takes text and returns a list of tokens.
type Tokenizer interface {
	Tokenize(text string) []Token
}

// normalize cleans the text: lowercases it and removes special characters.
func normalize(text string) string {
	// 1. Convert to lowercase
	text = strings.ToLower(text)

	// 2. Remove anything that isn't a letter or number
	// We use a builder for efficiency in Go
	var b strings.Builder
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// WordTokenizer splits text into words by spaces.
type WordTokenizer struct {
    Weight int // Importance of this tokenizer (usually high, e.g., 20)
}

func (t WordTokenizer) Tokenize(text string) []Token {
	cleanText := normalize(text)
	
	// Split by space
	words := strings.Fields(cleanText)
	
	var tokens []Token
	
	// Use a map to avoid duplicates (e.g. "go go go" becomes just one "go")
	seen := make(map[string]bool)

	for _, word := range words {
		// Filter out very short words like "a" or "I"
		if len(word) < 2 {
			continue
		}

		if !seen[word] {
			tokens = append(tokens, Token{
				Name:   word,
				Weight: t.Weight,
			})
			seen[word] = true
		}
	}

	return tokens
}