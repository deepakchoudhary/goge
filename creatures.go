package goge

import (
	"goge/extras"
)

type __Creature__ struct {
	Name          string
	Model         string
	Row           int
	Column        int
	__Direction__ string
	__Solidity__  bool
}

func (creature *__Creature__) __UpdateDirection__(dir string) *__Creature__ {
	if creature.__Direction__ != dir {
		creature.Model = extras.Flip(creature.Model)
		creature.__Direction__ = dir
	}

	return creature
}
func New(name string, syb string, solidity bool) *__Creature__ {
	return &__Creature__{Name: name, Model: syb, Row: 0, Column: 0, __Direction__: "R", __Solidity__: solidity}
}
