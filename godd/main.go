package main

import (
	"fmt"
	"godd/rpg"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	rand.Seed(time.Now().UnixNano())

	teo := rpg.NewPerson("Teo", "elfo", "clerigo")
	//fmt.Println("Personagem\n ", teo)

	nah := rpg.NewPerson("Nah", "humano", "ladino")
	//fmt.Println("Personagem\n ", nah)

	orc := rpg.NewCreature("Or1", "orc", "guerreiro")
	//fmt.Println("\n\nMonstro\n", orc)

	raid := rpg.NewGang("Raid", teo, nah) // time de players
	horde := rpg.NewGang("Horda", orc)    // time de monstros

	duel := rpg.NewDuel(*raid, *horde)

	fmt.Println(duel)

	duel.ExecDuel()

}
