package main

import (
	"fmt"
)

func main() {
	text := "Parser" // The word we are analyzing

	// 1. Exact Matcher
	wordTok := WordTokenizer{Weight: 20}
	fmt.Println("--- Word Tokenizer (Exact) ---")
	printTokens(wordTok.Tokenize(text))

	// 2. Autocomplete Matcher (Prefixes of length 3+)
	prefixTok := PrefixTokenizer{Weight: 5, MinLength: 3}
	fmt.Println("\n--- Prefix Tokenizer (Autocomplete) ---")
	printTokens(prefixTok.Tokenize(text))

	// 3. Typo Matcher (Trigrams / length 3)
	ngramTok := NGramTokenizer{Weight: 1, Length: 3}
	fmt.Println("\n--- N-Gram Tokenizer (Typos) ---")
	printTokens(ngramTok.Tokenize(text))
}

// Helper to print cleanly
func printTokens(tokens []Token) {
	for _, t := range tokens {
		fmt.Printf("['%s' : %d] ", t.Name, t.Weight)
	}
	fmt.Println()
}