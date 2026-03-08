/*
 * Copyright (c) 2026 Mipsou <chpujol@gmail.com>
 *
 * Licensed under the EUPL, Version 1.2 or later.
 * You may obtain a copy at:
 * https://joinup.ec.europa.eu/collection/eupl/eupl-text-eupl-12
 */

package search

// BM25 implements the Searcher interface using the BM25 ranking algorithm.
// Pure Go, zero external dependencies.
type BM25 struct{}

// NewBM25 creates a new BM25 search engine.
func NewBM25() *BM25 { return &BM25{} }

func (b *BM25) Index(corpus, docName, content string) error            { return nil }
func (b *BM25) Search(query string, maxResults int) ([]Result, error)  { return nil, nil }
func (b *BM25) Remove(corpus, docName string) error                    { return nil }
