package main

import (
	"fmt"
)

func main() {
	// Create a tokenizer with weight 20
	tokenizer := WordTokenizer{Weight: 20}

	text := "Go is fast, really fast!"
	tokens := tokenizer.Tokenize(text)

	fmt.Println("Original:", text)
	fmt.Println("Tokens found:")
	for _, t := range tokens {
		fmt.Printf("- '%s' (Weight: %d)\n", t.Name, t.Weight)
	}
};