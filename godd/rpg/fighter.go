package rpg

// A Fighter is any entity that can participate in a fight.
type Fighter interface {
	// CombatDice returns the resulting value for an attack of the Fighter
	// accounting for its attributes and attack bonuses.
	CombatDice() uint32

	// GetName returns the name of the Fighter.
	GetName() string

	// SubLife register damage on the Fighter account for its attributes and defense bonuses.
	SubLife(damage uint32)

	// GetLife returns the current life points of the Fighter
	GetLife() uint32
}
