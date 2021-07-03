package rpg

import "math/rand"

type Gang struct {
	Name     string
	Fighters []Fighter
}

func (g *Gang) SortOrder() {
	rand.Shuffle(len(g.Fighters), func(i, j int) { g.Fighters[i], g.Fighters[j] = g.Fighters[j], g.Fighters[i] })
}

func NewGang(name string, f ...Fighter) *Gang {
	gang := Gang{Fighters: f, Name: name}
	gang.SortOrder()
	return &gang
}
