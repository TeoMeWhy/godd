package rpg

import (
	"fmt"
	"godd/db"
	"godd/utils"
	"log"
	"math/rand"
	"os"
	"sort"
)

type Creature struct {
	Name             string
	Race             string
	Class            string
	Level            uint32
	Exp              uint32
	LifeDice         uint32
	TotalLife        uint32
	Attributes       map[string]uint32
	PrimaryAttribute string
	Modifiers        map[string]uint32
	Damage           uint32
	Armor            uint32
}

// NewCreature creates return a new creature based on name and race
func NewCreature(name, race, class string) *Creature {

	c := Creature{
		Name:  name,
		Race:  race,
		Class: class,
	}

	c.makePrimaryAttribute()
	c.createAttributes()
	c.makeLifeDice()
	c.calcModifiers()
	c.Damage = c.Modifiers[c.PrimaryAttribute]
	c.Armor = uint32(10) + c.Modifiers["destreza"]
	return &c
}

// MakePrimaryAttribute para atributos primários
func (c *Creature) makePrimaryAttribute() {

	att := map[string][]string{
		"clerigo":   {"sabedoria"},
		"guerreiro": {"forca", "destreza"},
		"ladino":    {"destreza"},
		"mago":      {"inteligencia"},
	}
	c.PrimaryAttribute = att[c.Class][rand.Intn(len(att[c.Class]))]
}

// RollDiceAttibute generates points to one attribute
func (c *Creature) rollDiceAttibute() int {

	dices := []int{}

	for len(dices) < 4 {
		dices = append(dices, (rand.Intn(6) + 1))
	}

	sort.Ints(dices)
	sum := 0

	for _, v := range dices[1:] {
		sum += v
	}
	return sum
}

// CreateAttributes makes a link between attributes and points by dices
func (c *Creature) createAttributes() {

	att := []string{
		"forca",
		"destreza",
		"constituicao",
		"inteligencia",
		"sabedoria",
		"carisma",
	}

	// order attributes by random excluding primaryAttribute
	newAtts := utils.RemoveElement(att, c.PrimaryAttribute)
	rand.Shuffle(len(newAtts), func(i, j int) { att[i], att[j] = att[j], att[i] })
	newAtts = append(newAtts, c.PrimaryAttribute)

	// generate all dices with same size of attributes
	points := []int{}
	for i := 0; i < len(newAtts); i++ {
		points = append(points, c.rollDiceAttibute())
	}
	sort.Ints(points)

	mapAtt := make(map[string]uint32)
	for i, v := range newAtts {
		mapAtt[v] = uint32(points[i])
	}

	c.Attributes = mapAtt
}

// CalcModifiers calculates all modifiers based on Attibutes
func (c *Creature) calcModifiers() {

	mods := make(map[string]uint32)
	raceMods := c.loadModifiers()

	for a, v := range c.Attributes {
		diff := (int(v) - 10) / 2

		if diff <= 0 {
			mods[a] = 0
		} else {
			mods[a] = uint32(diff)
		}

		mods[a] += raceMods[a]
	}
	c.Modifiers = mods
}

// MakeLifeDice calculates a Life Dice of creature
func (c *Creature) makeLifeDice() {

	data := map[string]uint32{
		"clerigo":   8,
		"guerreiro": 10,
		"ladino":    8,
		"mago":      6}

	c.LifeDice = data[c.Class] + c.Modifiers["constituicao"]
	c.TotalLife = c.LifeDice
}

func (c *Creature) String() string {

	txt := `
	Name:             %s
	Race:             %s
	Class:            %s
	Level:            %d
	Exp:              %d
	LifeDice:         %d
	PrimaryAttribute: %s
	Damage:           %d
	Armor:            %d
	Força:            %d
	Destreza:         %d
	Constituição:     %d
	Inteligência:     %d
	Sabedoria:        %d
	Carisma:          %d`

	return fmt.Sprintf(txt,
		c.Name,
		c.Race,
		c.Class,
		c.Level,
		c.Exp,
		c.LifeDice,
		c.PrimaryAttribute,
		c.Damage,
		c.Armor,
		c.Attributes["forca"],
		c.Attributes["destreza"],
		c.Attributes["constituicao"],
		c.Attributes["inteligencia"],
		c.Attributes["sabedoria"],
		c.Attributes["carisma"])
}

func (c *Creature) loadModifiers() map[string]uint32 {

	database := os.Getenv("DATABASE")
	fmt.Println(database)
	con := db.OpenSQLite(database)

	query := fmt.Sprintf(`
	SELECT forca,
		   destreza,
		   sabedoria,
		   inteligencia,
		   constituicao,
		   carisma
	FROM tb_raca
	where raca = '%s'
	limit 1;`, c.Race)

	rows, err := con.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var race string
	var forca, destreza, sabedoria, inteligencia, constituicao, carisma int

	data := make(map[string]uint32)

	for rows.Next() {
		err = rows.Scan(&race, &forca, &destreza, &sabedoria, &inteligencia, &constituicao, &carisma)
		if err != nil {
			panic("Erro ao capturar os dados da query")
		}

		data = map[string]uint32{
			"forca":        uint32(forca),
			"destreza":     uint32(destreza),
			"sabedoria":    uint32(sabedoria),
			"inteligencia": uint32(inteligencia),
			"constituicao": uint32(constituicao),
			"carisma":      uint32(carisma),
		}
	}
	return data
}

func (c *Creature) CombatDice() uint32 {
	return uint32(rand.Int31n(20)+1) + c.Modifiers[c.PrimaryAttribute]
}

func (c *Creature) GetName() string {
	return c.Name
}

func (c *Creature) SubLife(damage uint32) {

	if damage > c.TotalLife {
		c.TotalLife = 0
	} else {
		c.TotalLife -= damage
	}

}

func (c *Creature) GetLife() uint32 {
	return c.TotalLife
}
