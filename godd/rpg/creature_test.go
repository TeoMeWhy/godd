package rpg

import "testing"

func TestPrimAttribute(t *testing.T) {
	// FIXME: Refactor out the database dependency to allow isolated testing
	t.Skip("This test depends on database connection, skipping until refactoring out the dependency")
	creatures := []struct {
		creature *Creature
		expected string
	}{
		{
			creature: NewCreature("teo", "humano", "clerigo"),
			expected: "sabedoria",
		},
		{
			creature: NewCreature("teo", "humano", "ladino"),
			expected: "destreza",
		},
		{
			creature: NewCreature("teo", "humano", "mago"),
			expected: "inteligencia",
		},
	}

	for _, i := range creatures {
		res := i.creature.PrimaryAttribute
		if res != i.expected {
			t.Error("Esperado ", i.expected, ". Obtido ", res)
		}
	}
}
