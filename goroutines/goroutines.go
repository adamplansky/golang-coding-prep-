package goroutines

import "math/rand"

var alphabet = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

type IDGenerator struct {
	length int
	existing map[string]bool
	uniqueIDs chan string
}

func NewIDGenerator(length int) *IDGenerator {
	gen := &IDGenerator{length, map[string]bool{}, make(chan string)}
	go gen.generateInfiniteIDs()
	return gen
}

func (g *IDGenerator) createID() string {
	letters := make([]rune, g.length)
	for i:=0; i<g.length; i++ {
		letters[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(letters)
}


func (g *IDGenerator) generateInfiniteIDs() {
	for {
		id := g.createID()
		if ex, ok := g.existing[id]; !ok || !ex {
			g.existing[id] = true
			g.uniqueIDs <- id
		}
	}
}

func (g *IDGenerator) GetUniqueID() string {
	return <- g.uniqueIDs
}