package memory

import (
	"context"
	"sync"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"gonum.org/v1/gonum/mat"

	"github.com/go-aie/oneai/vectorstore/api"
)

// Memory is an in-memory vector store.
type Memory struct {
	documents map[string][]*api.Document
	mu        sync.RWMutex
}

func New() *Memory {
	return &Memory{
		documents: make(map[string][]*api.Document),
	}
}

func (m *Memory) Upsert(ctx context.Context, vendor string, documents []*api.Document) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, doc := range documents {
		sourceID := doc.Metadata.SourceID
		m.documents[sourceID] = append(m.documents[sourceID], doc)
	}
	return nil
}

func (m *Memory) Query(ctx context.Context, vendor string, vector []float64, topK int) ([]*api.Similarity, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if topK <= 0 {
		return nil, nil
	}

	target := mat.NewVecDense(len(vector), vector)

	similarities := make([]*api.Similarity, 0, topK) // Avoid null JSON array.
	for _, docs := range m.documents {
		for _, doc := range docs {
			candidate := mat.NewVecDense(len(doc.Vector), doc.Vector)
			score := mat.Dot(target, candidate)
			similarities = append(similarities, &api.Similarity{
				Document: doc,
				Score:    score,
			})
		}
	}

	// Sort similarities by score in descending order.
	slices.SortStableFunc(similarities, func(a, b *api.Similarity) int {
		if a.Score > b.Score {
			return -1
		} else if a.Score == b.Score {
			return 0
		} else {
			return 1
		}
	})

	if len(similarities) <= topK {
		return similarities, nil
	}
	return similarities[:topK], nil
}

// Delete deletes the chunks belonging to the given sourceIDs.
// As a special case, empty documentIDs means deleting all chunks.
func (m *Memory) Delete(ctx context.Context, vendor string, sourceIDs ...string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(sourceIDs) == 0 {
		maps.Clear(m.documents)
	}
	for _, sourceID := range sourceIDs {
		delete(m.documents, sourceID)
	}

	return nil
}
