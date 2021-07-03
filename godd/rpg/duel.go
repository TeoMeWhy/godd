package rpg

import (
	"fmt"
	"math/rand"
	"time"
)

type Duel struct {
	raid  Gang
	horde Gang
}

func (d *Duel) ExecMonstersDuel() {

	for len(d.raid.Fighters) > 0 && len(d.horde.Fighters) > 0 {

		d.RaidTurn()
		d.HordeTurn()
	}

}

func (d *Duel) HordeTurn() {
	fmt.Println("\nTurno da", d.horde.Name)
	nFighters := len(d.raid.Fighters)
	for _, f := range d.horde.Fighters {
		target := d.raid.Fighters[rand.Intn(nFighters)+1]
		d.ExecTurn(f, target)
	}
}

func (d *Duel) RaidTurn() {

	//TODO implementar o turno da Raid
	fmt.Println("\nTurno da", d.raid.Name)

}

func (d *Duel) ExecTurnAsync(attacker Fighter, enemyGang *Gang, ch chan string) {

	fmt.Println("Turno Executado por:", attacker.GetName())
	timeout := time.After(time.Second * 5) // tempo para escolha

	for {
		select {
		case target := <-ch:
			for _, i := range enemyGang.Fighters {
				if i.GetName() == target {
					fmt.Println(attacker.GetName(), "atacou", i)
					d.ExecTurn(attacker, i)
					return
				}
			}
			fmt.Println("Escola um target válido.")

		case <-timeout:
			fmt.Println("Nenhum alvo selecionado a tempo.")
			return
		}
	}
}

func (d *Duel) ExecTurn(attacker Fighter, target Fighter) {

	//TODO: implementar racional do dano e armadura
	target.SubLife(attacker.CombatDice())

}

func (d *Duel) String() string {

	raiders := ""
	for _, v := range d.raid.Fighters {
		raiders += v.GetName() + "\n"
	}

	horde := ""
	for _, v := range d.horde.Fighters {
		horde += v.GetName() + "\n"
	}

	txt := `
Raiders:
%s

Horde:
%s
`
	return fmt.Sprintf(txt, raiders, horde)
}

func NewDuel(raid Gang, horde Gang) *Duel {
	d := &Duel{raid: raid, horde: horde}
	return d
}
