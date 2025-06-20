package entities

import (
	"testing"

	"github.com/menterline/oncurve/entities"
)

func TestGetSpellReturnsAllMana(t *testing.T) {
	manaCost := []entities.Mana{entities.Black, entities.Black, entities.Blue}
	mySpell := entities.Spell{Cost: manaCost, CardType: entities.Creature}
	got := mySpell.GetManaCost()
	if len(got) != 3 {
		t.Errorf("expected mana cost length 3, got %d", len(got))
	}
	if mySpell.GetCardType() != entities.Creature {
		t.Errorf("expected card type Creature, got %v", mySpell.GetCardType())
	}
}

func TestCanCastSpell(t *testing.T) {
	manaMap := map[entities.Mana]int{
		entities.Black: 2,
		entities.Blue:  1,
	}
	manaPool := entities.ManaPool{AvailableMana: manaMap}

	spell := entities.Spell{
		Cost:     []entities.Mana{entities.Black, entities.Black, entities.Blue},
		CardType: entities.Creature,
	}

	if !spell.CanCast(manaPool) {
		t.Errorf("expected spell to be castable with the given mana pool")
	}

	spell.Cost = []entities.Mana{entities.Black, entities.Blue, entities.Red}
	if spell.CanCast(manaPool) {
		t.Errorf("expected spell not to be castable with the given mana pool")
	}
}
