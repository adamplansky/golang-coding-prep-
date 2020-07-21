package goroutines

import (
	"math/rand"
	"sync"
)


type MutexIDGenerator struct {
	length   int
	existing map[string]bool
	mux      sync.Mutex
}

func MutexNewIDGenerator(length int) *MutexIDGenerator {
	gen := &MutexIDGenerator{length: length, existing: map[string]bool{},
	}
	return gen
}

func (g *MutexIDGenerator) createID() string {
	letters := make([]rune, g.length)
	for i := 0; i < g.length; i++ {
		letters[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(letters)
}

func (g *MutexIDGenerator) generateInfiniteIDs() string {
	for {
		id := g.createID()
		if ex, ok := g.existing[id]; !ok || !ex {
			g.existing[id] = true
			return id
		}
	}
}

func (g *MutexIDGenerator) GetUniqueID() string {
	g.mux.Lock()
	id := g.generateInfiniteIDs()
	g.mux.Unlock()
	return id
}
