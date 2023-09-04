package task_1

import "math/rand"

type Boxer struct {
	Name           string
	Height, Weight int
	FistDamage     int
	HitPoints      int
	ChanceToEvade  int
}

func (b *Boxer) GetName() string {
	return b.Name
}

func (b *Boxer) Hit() int {
	return b.FistDamage
}

func (b *Boxer) CanEvade() bool {
	return rand.Intn(100) < b.ChanceToEvade
}

func (b *Boxer) IsAlive() bool {
	return b.HitPoints > 0
}

func (b *Boxer) TakeDamage(damage int) {
	b.HitPoints -= damage
}

func (b *Boxer) HpRemains() int {
	return b.HitPoints
}
