package rpg

type Person struct {
	Creature
	Capacity    float64
	CapacityMax float64
	Items       map[string]map[string]uint32
}

// NewPerson creates a new person using name, race and class
func NewPerson(name, race, class string) *Person {
	p := Person{Creature: *NewCreature(name, race, class)}
	p.calcCapacityMax()
	return &p
}

// calcCapacity calculates how much weight the person can carry
func (p *Person) calcCapacityMax() {
	p.CapacityMax = float64(p.Attributes["forca"]) * 7.5
}
