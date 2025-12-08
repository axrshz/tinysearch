package main

// Token represents a piece of text (like a word) and its importance.
// This mirrors the 'index_tokens' table in the blog.
type Token struct {
    ID     int    // Unique ID for the token
    Name   string // The actual text (e.g., "parser", "par")
    Weight int    // How important is this token type? (e.g., Exact word = 20, Prefix = 5)
}

// IndexEntry links a Token to a specific Document.
// This mirrors the 'index_entries' table in the blog.
type IndexEntry struct {
    ID           int // Unique ID for this entry
    TokenID      int // Which token is this? (Links to Token.ID)
    DocumentID   int // Which document acts this belong to?
    DocumentType int // Is this a Blog Post? A User Profile? A Product?
    FieldID      int // Where was it found? (Title? Body? Tags?)
    Weight       int // The final calculated score for this specific match
}

// Document is what we are actually searching for.
// In a real app, this might be a 'Post' or 'Product'.
type Document struct {
    ID      int
    Title   string
    Content string 
}