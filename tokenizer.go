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

// PrefixTokenizer generates the beginning parts of words.
type PrefixTokenizer struct {
	Weight    int
	MinLength int // Don't index prefixes shorter than this (e.g. "a", "ap")
}

func (t PrefixTokenizer) Tokenize(text string) []Token {
	cleanText := normalize(text)
	words := strings.Fields(cleanText)
	
	var tokens []Token
	seen := make(map[string]bool)

	for _, word := range words {
		// Crucial Step: Convert string to 'runes'.
		// In Go, strings are bytes. 'Runes' are actual characters.
		// If we don't do this, we might cut an emoji or accent mark in half!
		chars := []rune(word)
		
		if len(chars) < t.MinLength {
			continue
		}

		// Loop: create prefixes from MinLength up to the full word
		for i := t.MinLength; i <= len(chars); i++ {
			prefix := string(chars[:i])
			
			if !seen[prefix] {
				tokens = append(tokens, Token{
					Name:   prefix,
					Weight: t.Weight,
				})
				seen[prefix] = true
			}
		}
	}
	return tokens
}

// NGramTokenizer creates sliding windows of characters.
type NGramTokenizer struct {
	Weight int
	Length int // The size of the window (usually 3)
}

func (t NGramTokenizer) Tokenize(text string) []Token {
	cleanText := normalize(text)
	words := strings.Fields(cleanText)
	
	var tokens []Token
	seen := make(map[string]bool)

	for _, word := range words {
		chars := []rune(word)
		
		if len(chars) < t.Length {
			continue
		}

		// Slide the window across the word
		// If word is "Apple" (5 chars) and Length is 3:
		// i=0: "App"
		// i=1: "ppl"
		// i=2: "ple"
		for i := 0; i <= len(chars)-t.Length; i++ {
			gram := string(chars[i : i+t.Length])
			
			if !seen[gram] {
				tokens = append(tokens, Token{
					Name:   gram,
					Weight: t.Weight,
				})
				seen[gram] = true
			}
		}
	}
	return tokens
}