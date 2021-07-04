package rpg

import "math/rand"

type Gang struct {
	Name     string
	Fighters []Fighter
}

// ShuffleFighters shuffle the fighters of a Gang.
func (g *Gang) ShuffleFighters() {
	rand.Shuffle(len(g.Fighters), func(i, j int) { g.Fighters[i], g.Fighters[j] = g.Fighters[j], g.Fighters[i] })
}

// NewGang makes a new Gang with the given Fighters in a random order.
func NewGang(name string, f ...Fighter) *Gang {
	gang := Gang{Fighters: f, Name: name}
	gang.ShuffleFighters()
	return &gang
}
