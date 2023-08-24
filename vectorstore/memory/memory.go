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
	chunks map[string][]*api.Chunk
	mu     sync.RWMutex
}

func New() *Memory {
	return &Memory{
		chunks: make(map[string][]*api.Chunk),
	}
}

func (m *Memory) Upsert(ctx context.Context, vendor string, chunks map[string][]*api.Chunk) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for documentID, chunkList := range chunks {
		m.chunks[documentID] = append(m.chunks[documentID], chunkList...)
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

	var similarities []*api.Similarity
	for _, chunks := range m.chunks {
		for _, chunk := range chunks {
			candidate := mat.NewVecDense(len(chunk.Vector), chunk.Vector)
			score := mat.Dot(target, candidate)
			similarities = append(similarities, &api.Similarity{
				Chunk: chunk,
				Score: score,
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

// Delete deletes the chunks belonging to the given documentIDs.
// As a special case, empty documentIDs means deleting all chunks.
func (m *Memory) Delete(ctx context.Context, vendor string, documentIDs ...string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(documentIDs) == 0 {
		maps.Clear(m.chunks)
	}
	for _, documentID := range documentIDs {
		delete(m.chunks, documentID)
	}

	return nil
}
