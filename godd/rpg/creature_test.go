package rpg

import "testing"

type testprimaryatt struct {
	creature *Creature
	expected string
}

var creatures = []testprimaryatt{
	{creature: NewCreature("teo", "humano", "clerigo"),
		expected: "sabedoria"},

	{creature: NewCreature("teo", "humano", "ladino"),
		expected: "destreza"},

	{creature: NewCreature("teo", "humano", "mago"),
		expected: "inteligencia"},
}

func TestPrimAttribute(t *testing.T) {
	for _, i := range creatures {
		res := i.creature.PrimaryAttribute
		if res != i.expected {
			t.Error("Esperado ", i.expected, ". Obtido ", res)
		}
	}
}
