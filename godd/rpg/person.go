package rpg

type Person struct {
	Creature
	Capacity float64
	Items    map[string]map[string]uint32
}

// NewPerson creates a new person using name, race and class.
func NewPerson(name, race, class string) *Person {
	p := Person{
		Creature: *NewCreature(name, race, class),
	}
	return &p
}

// CapacityMax calculates how much weight the person can carry.
func (p *Person) CapacityMax() float64 {
	return float64(p.Attributes["forca"]) * 7.5
}
