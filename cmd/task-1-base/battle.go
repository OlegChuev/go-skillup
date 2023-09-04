package main

import (
	"fmt"
	"math/rand"

	"main.go/pkg/task_1"
)

type Battler interface {
	GetName() string
	Hit() int
	CanEvade() bool
	IsAlive() bool
	TakeDamage(int)
	HpRemains() int
}

func ChoseFighter(b *task_1.Boxer, k *task_1.Karateka) (attacker, defender Battler) {
	if rand.Intn(2) == 0 {
		attacker, defender = b, k
	} else {
		attacker, defender = k, b
	}

	return
}

func PrepareFighters() (b task_1.Boxer, k task_1.Karateka) {
	b = task_1.Boxer{Name: "Vasia", Height: 200, Weight: 120, FistDamage: 10, HitPoints: 113, ChanceToEvade: 15}
	k = task_1.Karateka{Name: "Kolia", Height: 175, Weight: 65, FistDamage: 8, FootDamage: 25, HitPoints: 82, ChanceToEvade: 25}

	return
}

// This is the duration of the fight to protect yourself from eternal random evades.
const FightDuration = 30

func main() {
	boxer, karateka := PrepareFighters()
	var attacker, defender Battler

	turn := 0

	for turn < FightDuration {
		attacker, defender = ChoseFighter(&boxer, &karateka)

		if defender.CanEvade() {
			fmt.Printf("%s evades the attack from %s.\n", defender.GetName(), attacker.GetName())
		} else {
			damage := attacker.Hit()
			defender.TakeDamage(damage)
			fmt.Printf("%s attacks %s for %d damage. %s has %d HP left.\n", attacker.GetName(), defender.GetName(), damage, defender.GetName(), defender.HpRemains())
		}

		// Check if the fight is over
		if !attacker.IsAlive() || !defender.IsAlive() {
			break
		}

		turn++
	}

	// Determine the winner
	switch {
	case attacker.IsAlive():
		fmt.Printf("%s wins the fight!\n", attacker.GetName())
	case defender.IsAlive():
		fmt.Printf("%s wins the fight!\n", defender.GetName())
	default:
		fmt.Println("It's a draw!")
	}
}
