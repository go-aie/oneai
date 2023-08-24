package api

import (
	"context"
)

//go:generate kungen -out=../controller -flat=false ./api.go VectorStore

type VectorStore interface {
	//kun:op POST /upsert
	//kun:param __ in=header name=Authorization required=true
	Upsert(ctx context.Context, vendor string, chunks map[string][]*Chunk) error

	//kun:op POST /query
	//kun:param __ in=header name=Authorization required=true
	Query(ctx context.Context, vendor string, vector []float64, topK int) (similarities []*Similarity, err error)

	//kun:op POST /delete
	//kun:param __ in=header name=Authorization required=true
	Delete(ctx context.Context, vendor string, documentIDs ...string) error
}

type Metadata struct {
	CorpusID string `json:"corpus_id,omitempty"`
}

type Document struct {
	ID       string   `json:"id,omitempty"`
	Text     string   `json:"text,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}

type Chunk struct {
	ID         string    `json:"id,omitempty"`
	Text       string    `json:"text,omitempty"`
	DocumentID string    `json:"document_id,omitempty"`
	Metadata   Metadata  `json:"metadata,omitempty"`
	Vector     []float64 `json:"vector,omitempty"`
}

type Similarity struct {
	*Chunk

	Score float64 `json:"score,omitempty"`
}
