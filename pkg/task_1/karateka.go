package task_1

import "math/rand"

type Karateka struct {
	Name           string
	Height, Weight int
	FistDamage     int
	FootDamage     int
	HitPoints      int
	ChanceToEvade  int
}

func (k *Karateka) GetName() string {
	return k.Name
}

func (k *Karateka) Hit() (hitDamage int) {
	if rand.Intn(2) == 0 {
		hitDamage = k.FistDamage
	} else {
		hitDamage = k.FootDamage
	}

	return
}

func (k *Karateka) CanEvade() bool {
	return rand.Intn(100) < k.ChanceToEvade
}

func (k *Karateka) IsAlive() bool {
	return k.HitPoints > 0
}

func (k *Karateka) TakeDamage(damage int) {
	k.HitPoints -= damage
}

func (k *Karateka) HpRemains() int {
	return k.HitPoints
}
