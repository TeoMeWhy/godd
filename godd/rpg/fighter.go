package rpg

type Fighter interface {
	CombatDice() uint32
	GetName() string
	SubLife(damage uint32)
	GetLife() uint32
}
