-- This stores the unique tokens
CREATE TABLE index_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    weight INTEGER NOT NULL,
    UNIQUE(name, weight)
);

-- This links tokens to documents
CREATE TABLE index_entries (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    token_id INTEGER NOT NULL,
    document_type INTEGER NOT NULL,
    document_id INTEGER NOT NULL,
    field_id INTEGER NOT NULL,
    weight INTEGER NOT NULL
);

-- Indexes make looking things up fast
CREATE INDEX idx_token_name ON index_tokens(name);
CREATE INDEX idx_entry_token ON index_entries(token_id);
CREATE INDEX idx_entry_doc ON index_entries(document_type, document_id);

-- THIS IS OUR "MAIN APP" DATA
-- The search engine indexes this, but doesn't "own" it.
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL
);