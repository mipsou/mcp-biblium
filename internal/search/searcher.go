/*
 * Copyright (c) 2026 Mipsou <chpujol@gmail.com>
 *
 * Licensed under the EUPL, Version 1.2 or later.
 * You may obtain a copy at:
 * https://joinup.ec.europa.eu/collection/eupl/eupl-text-eupl-12
 */

package search

// Result represents a single search result.
type Result struct {
	Collection  string  `json:"collection"`
	DocName string  `json:"doc_name"`
	Score   float64 `json:"score"`
	Snippet string  `json:"snippet"`
}

// Searcher is the interface for all search backends.
type Searcher interface {
	// Index processes a document and adds it to the search index.
	Index(collection, docName, content string) error

	// Search returns ranked results for a query.
	Search(query string, maxResults int) ([]Result, error)

	// Remove removes a document from the index.
	Remove(collection, docName string) error
}
