package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
	"godd/rpg"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Could not load dotenv files: %+v", err)
	}

	rand.Seed(time.Now().UnixNano())

	teo := rpg.NewPerson("Teo", "elfo", "clerigo")
	// fmt.Println("Personagem\n ", teo)

	nah := rpg.NewPerson("Nah", "humano", "ladino")
	// fmt.Println("Personagem\n ", nah)

	orc := rpg.NewCreature("Or1", "orc", "guerreiro")
	// fmt.Println("\n\nMonstro\n", orc)

	raid := rpg.NewGang("Raid", teo, nah) // time de players
	horde := rpg.NewGang("Horda", orc)    // time de monstros

	duel := rpg.NewDuel(*raid, *horde)

	fmt.Println(duel)

	duel.ExecMonstersDuel()
}
